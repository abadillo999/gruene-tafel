#!/bin/bash

DEPLOYMENT_NAME="gruene-taffel-deployment"
NAMESPACE="gruene_taffel"
IMAGE_REPOSITORY="dockerhub.com/badibadi"

_help() {
cat << EOF
Usage: ${0##*/} <args>
    -h, --help             Display help.
    -d, --deploy           Deploy.
    -c, --clean            Clean up deployment. 
EOF
}

_clean() {
    echo "CLEANING DEPLOYMENT\n"
    helm delete --purge ${DEPLOYMENT_NAME}
}

_deploy() {
    echo "HELM DEPENDENCY UPDATE\n"
    helm dep update helm/master
    [ $? -ne 0 ] && echo "Ops! Something went wrong..." && exit 1º
    # NOTE: Hostpath route "$PWD" coupled to GIT console on Windows in may not work on other O.S.
    
    echo "HELM DEPLOY\n"
    local command="helm install helm/master --name $DEPLOYMENT_NAME \
        --set ct.usvc.hostpath=/$PWD/test/integration-test/src \
        --set imagej-worker.usvc.hostpath=/$PWD/test/component-test/fiji \
        --debug"

    echo "Command to run: $command"
    output=$(${command} 2>&1)
    echo "Output: ${output}"

    [ $? -ne 0 ] && echo "Ops! Something went wrong..." && _clean && exit 1
}

[ $# -eq 0 ] && echo "No args, no party. Try ${0##*/} -h" && exit 0

while [ $# -ne 0 ]; do
    case "${1:-}" in
        -h|--help)
            _help
            exit
            ;;
        -d|--deploy)
            _deploy
            shift
            ;;
        -c|--clean)
            _clean
            shift
            ;;
        ?*)
            printf 'Unknown option \"%s\", try %s -h \n' "$1" "${0##*/}" >&2
            break
            ;;
    esac
done

exit