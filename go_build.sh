#!/bin/bash

# thiet lap che do strict mode de dung script neu co loi
set -e

# kiem tra xem co doi so dau vao (file main.go) khong
if [ -z "$1" ]; then
    echo "su dung: $0 <path/to/main.go>"
    exit 1
fi

input_file="$1"

# kiem tra xem file co ton tai khong
if [ ! -f "$input_file" ]; then
    echo "❌ loi: file '$input_file' khong ton tai!"
    exit 1
fi

# kiem tra go co duoc cai dat khong
if ! command -v go &> /dev/null; then
    echo "❌ loi: 'go' chua duoc cai dat hoac khong tim thay trong path."
    exit 1
fi

# ==============================================================================
# xac dinh duong dan va ten file
# ==============================================================================

# 1. xac dinh thu muc chua script nay (root project)
# day la noi script go_build.sh nam. du chay script tu dau, bien nay van tro ve dung cho.
script_dir="$( cd "$( dirname "${bash_source[0]}" )" &> /dev/null && pwd )"

# 2. lay duong dan tuyet doi cua file input
# neu input_file la tuong doi (khong bat dau bang /), ghep voi thu muc hien tai (pwd)
if [[ "$input_file" != /* ]]; then
    abs_input_file="$(pwd)/$input_file"
else
    abs_input_file="$input_file"
fi

# 3. lay ten thu muc chua file main.go de lam ten cho binary
# vi du: /path/to/algorithms/bubble_sort/main.go -> bubble_sort
source_dir=$(dirname "$abs_input_file")
base_name=$(basename "$source_dir")

# neu base_name la dau cham (.), nghia la nguoi dung dang dung ngay tai thu muc ma nguon
if [ "$base_name" == "." ]; then
    # lay ten thu muc cha thuc su
    base_name=$(basename "$(cd "$source_dir" && pwd)")
fi

# 4. tao thu muc output tai cung vi tri voi file source code
output_dir="${source_dir}/bin"
mkdir -p "$output_dir"

# don dep file cu trong thu muc output
rm -f "$output_dir"/*

# cac nen tang can build
# them linux/arm64 cho cac server arm/raspberry pi
platforms=("linux/amd64" "linux/arm64" "windows/amd64" "darwin/amd64" "darwin/arm64")

echo "🚀 bat dau build '$base_name' tu '$abs_input_file'..."
echo "📂 output directory: $output_dir"

# ham build de chay song song
build_platform() {
    local platform=$1
    local input=$2
    local out_dir=$3
    local name=$4

    # tach os va arch
    local target_os=${platform%/*}
    local target_arch=${platform#*/}
    
    # dat ten file output
    local output_name="${out_dir}/${name}-${target_os}-${target_arch}"
    
    if [ "$target_os" == "windows" ]; then
        output_name+=".exe"
    else
        output_name+=".bin"
    fi

    echo "  🔨 dang build cho $target_os/$target_arch..."
    
    # Build trong subshell để đảm bảo biến môi trường được set chính xác
    (
        export CGO_ENABLED=0
        export GOOS="$target_os"
        export GOARCH="$target_arch"
        
        go build -ldflags="-s -w" -o "$output_name" "$input"
    )

    if [ $? -ne 0 ]; then
        echo "  ❌ loi khi build cho $target_os/$target_arch"
        return 1
    fi

    # neu co upx (cong cu nen executable), thuc hien nen de giam kich thuoc them 50-70%
    # luu y: khong nen file tren macos (darwin) vi co the gay loi chu ky so (code signing/gatekeeper).
    if command -v upx &> /dev/null && [ "$target_os" != "darwin" ]; then
        # chay upx o che do im lang (-q) va nen tot nhat (--best)
        # bo qua loi neu upx khong nen duoc (vi du file qua nho hoac loi format)
        upx -q --best "$output_name" > /dev/null 2>&1
    fi
}

# chay build song song (parallel build)
# su dung mang pids de luu tru process id cua cac job chay ngam
pids=()
for platform in "${platforms[@]}"; do
    build_platform "$platform" "$abs_input_file" "$output_dir" "$base_name" &
    pids+=($!)
done

# cho tat ca cac process con hoan thanh
fail=0
for pid in "${pids[@]}"; do
    wait "$pid" || fail=1
done

if [ $fail -eq 0 ]; then
    echo "✅ hoan tat! da build thanh cong cho tat ca cac nen tang."
    
    # tao file checksum sha256 de kiem tra tinh toan ven
    echo "🔒 dang tao checksums..."
    checksum_file="${base_name}-sha256sums.txt"
    if command -v sha256sum &> /dev/null; then
        (cd "$output_dir" && sha256sum * > "$checksum_file")
        # verify lai ngay lap tuc
        echo "🔎 verifying checksums..."
        (cd "$output_dir" && sha256sum -c "$checksum_file")
    elif command -v shasum &> /dev/null; then
        (cd "$output_dir" && shasum -a 256 * > "$checksum_file")
        # verify lai ngay lap tuc
        echo "🔎 verifying checksums..."
        (cd "$output_dir" && shasum -a 256 -c "$checksum_file")
    fi

    # hien thi ket qua
    ls -lh "$output_dir"
else
    echo "❌ co loi xay ra trong qua trinh build."
    exit 1
fi
