#!/bin/bash

# 设置相关路径
binding_path="bindings"
out_path="out"

# 移除旧的bindings文件夹，并创建新的
rm -rf ${binding_path}
mkdir -p ${binding_path}

# 遍历编译后的合约文件并生成绑定
for file in "${out_path}"/*.sol/*.json; do
    # 提取合约名
    base=$(basename "$file" .json)                  # MyContract.json -> MyContract
    file_name=$(echo "$base" | tr '[:upper:]' '[:lower:]') # MyContract -> mycontract
    
    # 设置输出文件路径
    out_file="${binding_path}/${file_name}.go"
    abi_file="${binding_path}/${file_name}.abi"
    
    # 从JSON文件中提取ABI并修复类型
    jq '.abi | map(if .type == "listen" then .type = "event" else . end)' "$file" > "$abi_file"
    
    # 使用abigen生成Go绑定
    abigen \
        --abi="${abi_file}" \
        --pkg="bindings" \
        --type="${base}" \
        --out="${out_file}"
    # 删除abi文件
    rm -f ${abi_file}
done
