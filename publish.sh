# 请注意-这是测试版本生成脚本对应dist-testing分支，切记
#git subtree add --prefix=docker origin-wx release
#git subtree pull --prefix=dist-testing origin dist-testing
git checkout master
git pull

echo "如果推送失败就打开下面两节"
# git subtree add --prefix docker origin-wx release --squash
git subtree pull --prefix docker origin-wx release --squash

cd docker

git commit -a -m'自动编译 docker'
cd ../
# git subtree push --prefix=subtree/build/testing origin dist-testing
# 网校的测试环境分支
# git subtree push --prefix=subtree/build/testing origin release
git subtree push --prefix docker origin-wx release
