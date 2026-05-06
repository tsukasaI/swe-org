#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <string.h>

#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 8080

int connect_to_server(const char *ip, int port) {
  int sock = socket(AF_INET, SOCK_STREAM, 0);
  if (sock < 0) {
    perror("socket failed");
    return -1;
  }
  struct sockaddr_in addr;
  addr.sin_family = AF_INET;
  addr.sin_port = htons(port);
  addr.sin_addr.s_addr = inet_addr(ip);
  if (connect(sock, (struct sockaddr *)&addr, sizeof(addr)) < 0) {
    perror("connect failed");
    close(sock);
    return -1;
  }
  return sock;
}

void send_request(int sock, const char *host, int port, const char *path) {
  char request[1024];
  int request_len = snprintf(request, sizeof(request), 
      "GET %s HTTP/1.1\r\n"
      "Host: %s:%d\r\n"
      "\r\n",
      path, host, port);

  write(sock, request, request_len);
}

ssize_t read_response(int sock, char *buf, size_t bufsize) {
  ssize_t total = 0;
  ssize_t n;
  while ((n = read(sock, buf + total, bufsize - total - 1)) > 0) {
    total += n;
  }
  if (n < 0) {
    return -1;
  }
  buf[total] = '\0';
  return total;
}

int main(void) {
  int sock = connect_to_server(SERVER_IP, SERVER_PORT);
  if (sock < 0) {
    exit(1);
  }

  send_request(sock, SERVER_IP, SERVER_PORT, "/calc?query=2+10");

  char response[4096];
  if (read_response(sock, response, sizeof(response)) < 0) {
    perror("read failed");
    exit(1);
  }

  char *header_end = strstr(response, "\r\n\r\n");

  if (header_end == NULL) {
    fprintf(stderr, "malformed response (no header terminator)\n");
    exit(1);
  }

  char *body = header_end + 4;
  printf("result: %s\n", body);

  close(sock);
  return 0;
}
