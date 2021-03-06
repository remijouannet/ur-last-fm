#!/bin/bash

set +e

version=$(git describe --abbrev=0 --tags)

which github-release || echo 'please install the tool github-releases'

github-release info \
    --user remijouannet \
    --repo ur-last-fm \
    --tag $version

if [ $? != 0 ]
then
    echo "the release doesn't exist"
    github-release release \
        --user remijouannet \
        --repo ur-last-fm \
        --tag $version \
        --pre-release \
        --name "$version" \
        --description "$version" || echo "failed to create release for $version"
fi

cd pkg/

rm -f *.zip
 
ls | while read binary
do
    echo "zipping $binary"
    zip -9 -r $binary.zip $binary
    echo "upload $binary"
    github-release upload \
        --user remijouannet \
        --name "$binary.zip" \
        --repo ur-last-fm \
        --file "$binary.zip" \
        --replace \
        --tag $version || echo "failed to upload $binary"
done

