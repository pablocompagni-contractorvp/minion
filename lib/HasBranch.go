package lib

import "fmt"

func HasBranch(path, branch string) (bool, error) {
	response, _, err := RunCommand("git", path, "show-ref", fmt.Sprintf("refs/heads/%v", branch))
	if err != nil {
		return false, nil
	}

	if len(response) > 0 {
		return true, nil
	}

	return false, nil
}
