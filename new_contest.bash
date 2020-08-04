#!/bin/bash
set -eu

if [ $# -lt 4 ]; then
    echo "specify the contest name, the contest nubmer, the contest division and the number of tasks"
    exit 1
fi

contest_name=$1
contest_number=$2
contest_division=$3
task_num=$4

lang="go"
if [ $# -eq 5 ]; then
    lang=$5
fi

cd $(dirname $0)
cwd=$(pwd)

division_dir="${cwd}/${contest_name}/${contest_number}/${contest_division}"
if [ -e ${division_dir} ]; then
    echo "${division_dir} directory already exists."
    exit 1
fi
mkdir -p ${division_dir}

tasks=("a" "b" "c" "d" "e" "f" "g" "h" "i" "j" "k" "l" "m" "n" "o" "p" "q" "r" "s" "t" "u" "v" "w" "x" "y" "z")
for i in ${!tasks[@]}; do
    if [ $i -ge $task_num ]; then
        break
    fi
    mkdir "${division_dir}/${tasks[$i]}"
    cp ${cwd}/template.${lang} ${division_dir}/${tasks[$i]}/main.${lang}
done
