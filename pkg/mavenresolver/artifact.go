package mavenresolver

import (
	"encoding/xml"
	"fmt"
	"github/rashedkvm/maven-artifact/pkg/repository"
	"io/ioutil"
	"net/http"
	"strings"
)

type Repository struct {
	URL      string
	Username string
	Password string
}

type Artifact struct {
	Id          string
	GroupId     string
	Version     string
	ResolvedURL string
	MetaXML     string
}

func (r *Artifact) Resolve(repo *Repository) error {

	if repo == nil {
		return fmt.Errorf("missing repo or artifact object")
	}

	groupIdSplit := strings.Split(r.GroupId, ".")

	var metadataURL = fmt.Sprintf("%s/%s/%s/%s/maven-metadata.xml", repo.URL, strings.Join(groupIdSplit, "/"), r.Id, r.Version)

	cl := repository.Client()

	req, err := http.NewRequest("GET", metadataURL, nil)
	if err != nil {
		return err
	}
	if repo.Username != "" && repo.Password != "" {
		req.SetBasicAuth(repo.Username, repo.Password)
	}

	resp, err := cl.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var metadata Metadata
	err = xml.Unmarshal(body, &metadata)
	if err != nil {
		return fmt.Errorf("error processing metadata %v \n %v", string(body), err)
	}

	fmt.Printf("url: %q \n status %q \n", metadataURL, resp.Status)

	if strings.HasSuffix(r.Version, "-SNAPSHOT") {
		snapshotVersions := Filter(metadata.Versioning.SnapshotVersions.SnapshotVersion, func(v SnapshotVersion) bool {
			return v.Extension == "jar"
		})

		if len(snapshotVersions) > 0 {
			snapshotMetaxml, err := xml.MarshalIndent(snapshotVersions, "  ", "    ")
			if err != nil {
				return err
			}
			r.MetaXML = string(snapshotMetaxml)

			// select snapshot version
			snapshotversion := fmt.Sprintf("%s-%s-%s-%s", r.Id, strings.Replace(r.Version, "-SNAPSHOT", "", 1), metadata.Versioning.Snapshot.Timestamp, metadata.Versioning.Snapshot.BuildNumber)
			r.ResolvedURL = fmt.Sprintf("%s/%s/%s/%s/%s.jar", repo.URL, strings.Join(groupIdSplit, "/"), r.Id, r.Version, snapshotversion)

			return nil
		}
	}

	metaxml, err := xml.MarshalIndent(metadata, "  ", "    ")
	if err != nil {
		return err
	}
	r.MetaXML = string(metaxml)

	return nil

}
