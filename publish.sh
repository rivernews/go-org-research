# Make sure to commit and push first

VERSION=${1:-v0.1.0}
git tag ${VERSION}
git push origin ${VERSION}
echo "Registering new version ${VERSION} to go..."
GOPROXY=proxy.golang.org go list -m github.com/rivernews/go-org-research@${VERSION}
