#!/bin/bash

NAMESPACE="gruene_taffel"
IMAGE_REPOSITORY="dockerhub.com/badibadi"

_help() {
cat << EOF
Usage: ${0##*/} <args>
    -h, --help             Display help
    -d, --deploy           Deploy
    -c, --clean            Clean up namespace 
EOF
}

main() {    
    local deploy=
    local clean=

    while :; do
        case "${1:-}" in
            -h|--help)
                _help
                exit
                ;;
            -d|--deploy)
                deploy=all
                ;;
            -c|--clean)
                clean=true
                ;;
            -?*)
                printf 'WARN: Unknown option (ignored): %s\n' "$1" >&2
                ;;
            *)
                break
                ;;
        esac

        shift
    done

    if [[ -n "$deploy" ]];
    then
        echo "HELM DEPENDENCY UPDATE\n"
        #helm dep update master/helm
        git status
        echo "HELM DEPLOY\n"
        git status
    fi

    if [[ -n "$clean" ]];
    then
        helm 
    fi

    popd
}

main "$@"