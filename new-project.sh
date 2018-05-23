#!/usr/bin/env bash
echo -e "\033[1;35m<< New FSM Project >>\033[1;0m";
echo -ne "\033[1;34mProject Package (ex: \033[1;35mgithub.com/BrandonRomano/project\033[1;34m): \033[1;0m"; read package

if [[ "$package" = "" ]]; then
    echo -e '\033[1;31mERROR: Package must be set\033[1;0m'
    exit 1
fi

# Ensure GOPATH is set properly
if [[ ! -d "$GOPATH" ]]; then
    echo -e '\033[1;31mERROR: $GOPATH is not set properly\033[1;0m'
    echo 'See https://github.com/golang/go/wiki/SettingGOPATH for more details'
    exit 1
fi

# Clone
echo -e "\033[1;35mCloning fsm/getting-started...\033[1;0m";
packagedir="$GOPATH/src/$package"
git clone https://github.com/fsm/getting-started.git "$packagedir"

# Clean up
echo -e "\033[1;35mCleaning up...\033[1;0m";
cd "$packagedir"
rm -rf .git
rm .gitignore
rm LICENSE.md
rm README.md
rm new-project.sh

# Replace the Package
find . | grep ".go" | grep -v vendor | while read file; do
    from="github.com\/fsm\/getting-started"
    to="$(echo "$package" | sed "s/\//\\\\\//g")"
    sed -i '' "s/$from/$to/g" "$file"
done

# OK
echo -e "\033[0;32mProject setup at '$packagedir'\033[1;0m"
