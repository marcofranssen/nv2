package registry

import (
	digest "github.com/opencontainers/go-digest"
	oci "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/notaryproject/notary/v2/signature"
)

func OCIDescriptorFromNotary(desc signature.Descriptor) oci.Descriptor {
	return oci.Descriptor{
		MediaType: desc.MediaType,
		Digest:    digest.Digest(desc.Digest),
		Size:      desc.Size,
	}
}
