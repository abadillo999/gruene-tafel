#!/bin/bash

DEPLOYMENT_NAME="gruene-taffel-deployment"
NAMESPACE="gruene_taffel"
IMAGE_REPOSITORY="dockerhub.com/badibadi"

_help() {
cat << EOF
Usage: ${0##*/} <args>
    -h, --help             Display help.
    -d, --deploy           Deploy.
    -c, --clean            Clean up namespace. 
EOF
}

_clean() {
    echo "CLEANING DEPLOYMENT\n"
    helm delete --purge ${DEPLOYMENT_NAME}
}

_deploy() {
    echo "HELM DEPENDENCY UPDATE\n"
    helm dep update helm/controller
    [ $? -ne 0 ] && echo "Ops! Something went wrong..." && exit 1

    echo "HELM DEPLOY\n"
    helm install helm/controller --name ${DEPLOYMENT_NAME}
    [ $? -ne 0 ] && echo "Ops! Something went wrong..." && _clean && exit 1
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
        _deploy
    fi

    if [[ -n "$clean" ]];
    then
        _clean
    fi

}

[ $# -eq 0 ] && echo "No args, no party, try ${0##*/} -h"
main "$@"
