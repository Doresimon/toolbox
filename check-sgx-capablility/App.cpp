/*
 * Usage::
 * make
 * LD_LIBRARY_PATH=$SGX_SDK/lib64 ./app
 */


#include <stdio.h>
#include <string>

#include "sgx_capable.h"

#include "App.h"


std::string parse_sgx_device_status(sgx_device_status_t s) {
    std::string ret;
    switch (s)
    {
    case SGX_ENABLED:
        ret = "sgx is enabled";
        break;
    case SGX_DISABLED_REBOOT_REQUIRED:
        ret = "A reboot is required to finish enabling SGX";
        break;
    case SGX_DISABLED_LEGACY_OS:
        ret = "SGX is disabled and a Software Control Interface is not available to enable it";
        break;
    case SGX_DISABLED:
        ret = "SGX is not enabled on this platform. More details are unavailable.";
        break;
    case SGX_DISABLED_SCI_AVAILABLE:
        ret = "SGX is disabled, but a Software Control Interface is available to enable it";
        break;
    case SGX_DISABLED_MANUAL_ENABLE:
        ret = "SGX is disabled, but can be enabled manually in the BIOS setup";
        break;
    case SGX_DISABLED_HYPERV_ENABLED:
        ret = "Detected an unsupported version of Windows* 10 with Hyper-V enabled";
        break;
    case SGX_DISABLED_UNSUPPORTED_CPU:
        ret = "SGX is not supported by this CPU";
        break;
    
    default:
        break;
    }

    return ret;
}

int check_sgx_cap_status() {
    printf("[~] get sgx capable status...\n");
    sgx_device_status_t ds;
    sgx_status_t ret = SGX_ERROR_UNEXPECTED;
    ret = sgx_cap_get_status(&ds);
    if (ret != SGX_SUCCESS) {
        printf("[-] error\n");
        printf("[-] error_code = 0x%.4x\n", ret);
        return -1;
    }

    printf("[+] device_status_code = 0x%.4x\n", ds);
    printf("[+] device_status_notice = %s\n", parse_sgx_device_status(ds).c_str());
    return 0;
}

int check_sgx_is_capable() {
    printf("[~] check if sgx is capable...\n");

    char const * CAP_EQ_0_NOTICE = "Intel SGX device is not available or may require manual configuration.";
    char const * CAP_EQ_1_NOTICE = "Platform is enabled for the Intel SGX or the Software Control Interface is available to configure the Intel SGX device.";

    sgx_status_t ret = SGX_ERROR_UNEXPECTED;

    int capable = -1;
    ret = sgx_is_capable(&capable);
    if (ret != SGX_SUCCESS) {
        printf("[-] error\n");
        printf("[-] error_code = 0x%.4x\n", ret);
        return -1;
    }

    if (capable==0) {
        printf("[-] %s\n", CAP_EQ_0_NOTICE);
    }else{
        printf("[+] %s\n", CAP_EQ_1_NOTICE);
    }

    return 0;
}

int enable_sgx_capable() {
    printf("[~] try to enable sgx capability...\n");
    sgx_device_status_t ds;
    sgx_status_t ret = SGX_ERROR_UNEXPECTED;

    ret = sgx_cap_enable_device(&ds);
    if (ret != SGX_SUCCESS) {
        printf("[-] error\n");
        printf("[-] error_code = 0x%.4x\n", ret);
        return -1;
    }

    printf("[+] device_status_code = 0x%.4x\n", ds);
    printf("[+] device_status_notice = %s\n", parse_sgx_device_status(ds).c_str());

    return 0;
}

/* Application entry */
int main(int argc, char *argv[])
{
    (void)(argc);
    (void)(argv);

    check_sgx_is_capable();

    check_sgx_cap_status();

    enable_sgx_capable();
    
    printf("[~] check finished.\n");

    return 0;
}

