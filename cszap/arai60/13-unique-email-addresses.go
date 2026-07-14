// https://leetcode.com/problems/unique-email-addresses/
func numUniqueEmails(emails []string) int {
    set := make(map[string]struct{})

    for _, v := range emails {
        splitted := strings.Split(v, "@")
        local := splitted[0]
        domain := splitted[1]
        splittedByPlus := strings.Split(local, "+")
        normalizedLocal := strings.ReplaceAll(splittedByPlus[0], ".", "")

        set[fmt.Sprintf("%s@%s", normalizedLocal, domain)] = struct{}{}
    }
    return len(set)
}
