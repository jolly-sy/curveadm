#!/usr/bin/env

g_version=$1
g_nos=${NOS}
g_root=$(pwd)/.build
g_curveadm=${g_root}/curveadm
g_curveadm_bin=${g_curveadm}/bin
rm -rf ${g_root}

mkdir -p ${g_curveadm_bin}
cp bin/curveadm ${g_curveadm_bin}
[[ -f .CHANGELOG ]] && cp .CHANGELOG ${g_curveadm}/CHANGELOG
(cd ${g_curveadm} && ./bin/curveadm -v && ls -ls bin/curveadm && [[ -f CHANGELOG ]] && cat CHANGELOG)
(cd ${g_root} && tar -zcf curveadm-${g_version}.tar.gz curveadm)

read -p "Do you want to upload curveadm-${g_version}.tar.gz to NOS? " input
case $input in
    [Yy]* )
        if [ -z ${g_nos} ]; then
            echo "nos: command not found"
            exit 1
        fi
        ${g_nos} -putfile \
            ${g_root}/curveadm-${g_version}.tar.gz \
            curveadm \
            -key release/curveadm-${g_version}.tar.gz \
            -replace true
        ;;
    [Nn]* )
        exit
        ;;
    * )
        echo "Please answer yes or no."
        ;;
esac
