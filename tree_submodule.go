package git

// GetSubmodules returns the submodules in this commit tree (parsed from .gitmodules)
func (t *Tree) GetSubmodules() ([]*Submodule, error) {
	blob, err := t.GetBlobByPath(".gitmodules")
	if err != nil {
		if err == ErrNotExist {
			// No submodules in this commit, that's cool.
			return nil, nil
		}
		return nil, err
	}

	r, err := blob.Data()
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return parseSubmoduleConfig(r)
}
