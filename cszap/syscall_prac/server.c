#include <stdio.h>
#include <errno.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <errno.h>
#include <signal.h>

#define PORT 8080
#define BUFFER_SIZE 1024

volatile sig_atomic_t should_stop = 0;

// 割り込みの制御
void signal_handler(int signum) {
  (void)signum;
  should_stop = 1;
}

// HTTP responseを生成
void send_response(int client_fd, int status, const char *status_text,
    const char *body, int body_len) {
    char response[1024];
    int response_len = snprintf(response, sizeof(response),
        "HTTP/1.1 %d %s\r\n"
        "Content-Type: text/plain\r\n"
        "Content-Length: %d\r\n"
        "\r\n"
        "%s", status, status_text, body_len, body);
    write(client_fd, response, response_len);
}

// 文字列を "+" で区切って数値として足し算をする
int calc_sum(char *value) {
    int sum = 0;
    char *token = strtok(value, "+");
    while (token != NULL) {
      sum += atoi(token);
      token = strtok(NULL, "+");
    }
    return sum;
}

// ハンドラ
void handle_request(int client_fd) {
  char buffer[BUFFER_SIZE] = {0};
  ssize_t n = read(client_fd, buffer, sizeof(buffer) - 1);
  if (n < 0) {
    perror("read failed");
    close(client_fd);
    return;
  }
  printf("received %zd bytes:\n%s\n", n, buffer);

  char method[16] = {0};
  char path[256] = {0};
  sscanf(buffer, "%15s %255s", method, path);
  printf("parsed request:\nmethod = %s\npath=%s\n", method, path);

  char *query = NULL;
  char *q = strchr(path, '?');
  if (q != NULL) {
    *q = '\0';
    query = q + 1;
  }

  printf("path: %s\nquery=%s\n", path, query ? query : "(none)");

  char body[256];
  int body_len;
  int status;
  const char *status_text;

  if (strcmp(path, "/calc") != 0) {
    // 404
    status = 404;
    status_text = "Not Found";
    body_len = snprintf(body, sizeof(body), "Not Found");
  } else if (query == NULL || strncmp(query, "query=", 6) != 0) {
    // 400
    status = 400;
    status_text = "Bad Request";
    body_len = snprintf(body, sizeof(body), "Bad Request");
  } else {
    // 200
    char *value = query + 6;
    int sum = calc_sum(value);
    status = 200;
    status_text = "OK";
    body_len = snprintf(body, sizeof(body), "%d", sum);
  }

  send_response(client_fd, status, status_text, body, body_len);
  close(client_fd);
}

int main() {
  struct sigaction sa = {0};
  sa.sa_handler = signal_handler;
  if (sigaction(SIGINT, &sa, NULL) < 0) {
    perror("sigaction failed");
    exit(1);
  }

  int server_fd = socket(AF_INET, SOCK_STREAM, 0);
  if (server_fd == -1) {
    perror("socket failed");
    exit(1);
  }
  printf("server_fd: %d\n", server_fd);

  struct sockaddr_in address;
  address.sin_family = AF_INET;
  address.sin_addr.s_addr = htonl(INADDR_ANY);
  address.sin_port = htons(PORT);

  int opt = 1;
  setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));
  if (bind(server_fd, (struct sockaddr *)&address, sizeof(address))) {
    perror("bind failed");
    exit(1);
  }

  if (listen(server_fd, 5)) {
    perror("listen failed");
    exit(1);
  }

  printf("listening on port: %d\n", PORT);

  struct sockaddr_in client_addr;

  while (!should_stop) {
    socklen_t addrlen = sizeof(client_addr);
    int client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &addrlen);
    if (client_fd < 0) {
      if (errno == EINTR) {
        continue;
      }
      perror("accept failed");
      continue;
    }
    handle_request(client_fd);
  }

  printf("\nshutting down\n");
  close(server_fd);
  return 0;
}
