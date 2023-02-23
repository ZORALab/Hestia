echo >/dev/null # >nul & GOTO WINDOWS & rem ^
################################################################################
# Linux Variables                                                              #
################################################################################
pwd="$PWD"
################################################################################
# Linux Library Codes                                                          #
################################################################################
_print_status() {
        __status_message=""
        case "$1" in
        warn|Warning|WARNING|warning)
                __status_message="[ WARNING ]"
                shift 1
                ;;
        err|Error|ERROR|error)
                __status_message="[  ERROR  ]"
                shift 1
                ;;
        info|Info|INFO)
                __status_message="[  INFO   ]"
                shift 1
                ;;
        ok|Ok|OK)
                __status_message="[   OK    ]\n"
                shift 1
                ;;
        *)
                __status_message="[  INFO   ]"
                ;;
        esac

        1>&2 printf "%s %s\n" "$__status_message" "$@"
        unset __status_message
}

_exit_main() {
        case $1 in
        0)
                exit 0
                ;;
        *)
                exit 1
                ;;
        esac
}

################################################################################
# Linux Library Codes                                                          #
################################################################################
################################################################################
# Linux Main Codes                                                             #
################################################################################
check_hugo_availability() {
        if [ "$(type -p 'hugo')" == "" ]; then
                _print_status error "missing hugo program."
                _exit_main 1
        fi
}




change_directory_to_hugo_config() {
        cd "${BASH_SOURCE%/*}"
}




return_to_current_directory() {
        cd "$PWD"
        printf "\n"
}




run_hugo() {
        hugo server --buildDrafts \
                --noBuildLock \
                --disableFastRender \
                --bind "localhost" \
                --baseURL "http://localhost" \
                --port 8080 \
                --cleanDestinationDir \
                --gc
}




main() {
        check_hugo_availability
        change_directory_to_hugo_config
        run_hugo
        return_to_current_directory
        _exit_main 0
}
main $@
################################################################################
# Linux Main Codes                                                             #
################################################################################




:WINDOWS
::##############################################################################
:: Windows Main Codes                                                          #
::##############################################################################
hugo server --buildDrafts ^
        --noBuildLock ^
        --disableFastRender ^
        --bind "localhost" ^
        --baseURL "http://localhost" ^
        --port 8080 ^
        --cleanDestinationDir ^
        --gc
::##############################################################################
:: Windows Main Codes                                                          #
::##############################################################################
