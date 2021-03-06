package openshift

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	privateClusterOutboundFilename = "private-cluster-outbound-service.yaml"
)

var _ asset.WritableAsset = (*PrivateClusterOutbound)(nil)

// PrivateClusterOutbound generates the private-cluster-outbound-*.yml files
type PrivateClusterOutbound struct {
	FileList []*asset.File
}

// Name returns a human friendly name for the asset.
func (*PrivateClusterOutbound) Name() string {
	return "Private Cluster Outbound Service"
}

// Dependencies returns all of the dependencies directly needed by the asset
func (*PrivateClusterOutbound) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Generate generates the actual files by this asset
func (p *PrivateClusterOutbound) Generate(dependencies asset.Parents) error {
	data, err := content.GetOpenshiftTemplate(privateClusterOutboundFilename)
	if err != nil {
		return err
	}

	p.FileList = append(p.FileList, &asset.File{
		Filename: filepath.Join(content.TemplateDir, privateClusterOutboundFilename),
		Data:     []byte(data),
	})
	return nil
}

// Files returns the files generated by the asset.
func (p *PrivateClusterOutbound) Files() []*asset.File {
	return p.FileList
}

// Load returns the asset from disk
func (p *PrivateClusterOutbound) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, privateClusterOutboundFilename))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	p.FileList = append(p.FileList, file)

	return true, nil
}
