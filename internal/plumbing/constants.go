package plumbing

// Repository-level directory & file names used by plumbing layer.
// These follow Git's structure: .minigit/{objects, refs/heads, HEAD}
const (
	RepoDir         = ".minigit"
	ObjectsDir      = "objects"
	RefsDir         = "refs"
	HeadsDir        = "heads"
	HeadFile        = "HEAD"
	DefaultBranch   = "main"
	HeadFileContent = "ref: refs/heads/" + DefaultBranch + "\n"
)
