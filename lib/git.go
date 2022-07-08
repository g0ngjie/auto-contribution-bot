package lib

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// 本地仓库目录
var gitDir = "git-repo"

// git clone
func gitClone() {
	// 判断本地是否存在git仓库
	// 如果不存在则拉取
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		// 获取当前项目根路径
		url, directory, username, password := Cfg.GitUrl, gitDir, Cfg.GitUser, Cfg.GitPass

		_, err := git.PlainClone(directory, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: username,
				Password: password,
			},
			URL:      url,
			Progress: os.Stdout,
		})
		if err != nil {
			logrus.Errorf("拉取仓库异常：%s", err.Error())
		}
	}
}

// git pull
func gitPull(directory string) (*git.Worktree, *git.Repository) {
	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	if err != nil {
		logrus.Errorf("打开仓库异常：%s", err.Error())
	}

	w, err := r.Worktree()
	if err != nil {
		logrus.Errorf("git pull 异常：%s", err.Error())
	}

	w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: Cfg.GitUser,
			Password: Cfg.GitPass,
		},
	})
	return w, r
}

// git commit
func gitCommit(w *git.Worktree, content *string) {
	// Adds the new file to the staging area.
	// Info("git add example-git-file")
	_, err := w.Add(".")
	if err != nil {
		logrus.Errorf("git add 异常：%s", err.Error())
	}

	// We can verify the current status of the worktree using the method Status.
	// Info("git status --porcelain")
	status, err := w.Status()
	if err != nil {
		logrus.Errorf("git status 异常：%s", err.Error())
	}

	logrus.Infof("%v", status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	// Info("git commit -m \"example go-git commit\"")
	_, err = w.Commit(*content, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "bot",
			Email: Cfg.GitEmail,
			When:  time.Now(),
		},
	})

	if err != nil {
		logrus.Errorf("git commit 异常：%s", err.Error())
	}
}

// git push
func gitPush(r *git.Repository) {
	err := r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: Cfg.GitUser,
			Password: Cfg.GitPass,
		},
		Progress: os.Stdout,
	})
	if err != nil {
		logrus.Errorf("git push 异常：%s", err.Error())
	}
}

func BotRun() {
	// 获取当前项目根路径
	rootPath, _ := os.Getwd()
	directory := filepath.Join(rootPath, gitDir)
	// 仓库初始化
	gitClone()
	worktree, repository := gitPull(directory)
	// 写文件
	content := writeFile(gitDir)
	// 提交
	gitCommit(worktree, content)
	// 推送
	gitPush(repository)

}
