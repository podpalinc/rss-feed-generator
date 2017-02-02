package podcast

import (
	"encoding/xml"
	"time"
)

// Item represents a single entry in a podcast.
//
// Article minimal requirements are:
// - Title
// - Description
// - Link
//
// Audio minimal requirements are:
// - Title
// - Description
// - Enclosure (HREF, Type and Length all required)
//
// Recommendations:
// - Setting the minimal fields sets most of other fields, including iTunes.
// - Use the Published time.Time setting instead of PubDate.
// - Always set an Enclosure.Length, to be nice to your downloaders.
// - Use Enclosure.Type instead of setting TypeFormatted for valid extensions.
type Item struct {
	XMLName          xml.Name   `xml:"item"`
	GUID             string     `xml:"guid"`
	Title            string     `xml:"title"`
	Link             string     `xml:"link"`
	Description      string     `xml:"description"`
	Author           *Author    `xml:"-"`
	AuthorFormatted  string     `xml:"author,omitempty"`
	Category         string     `xml:"category,omitempty"`
	Comments         string     `xml:"comments,omitempty"`
	Source           string     `xml:"source,omitempty"`
	PubDate          *time.Time `xml:"-"`
	PubDateFormatted string     `xml:"pubDate,omitempty"`
	Enclosure        *Enclosure

	// https://help.apple.com/itc/podcasts_connect/#/itcb54353390
	IAuthor   string `xml:"itunes:author,omitempty"`
	ISubtitle string `xml:"itunes:subtitle,omitempty"`
	// TODO: CDATA
	ISummary           string `xml:"itunes:summary,omitempty"`
	IImage             *IImage
	IDuration          string `xml:"itunes:duration,omitempty"`
	IExplicit          string `xml:"itunes:explicit,omitempty"`
	IIsClosedCaptioned string `xml:"itunes:isClosedCaptioned,omitempty"`
	IOrder             string `xml:"itunes:order,omitempty"`
}

// AddEnclosure adds the downloadable asset to the podcast Item.
func (i *Item) AddEnclosure(
	url string, enclosureType EnclosureType, lengthInSeconds int64) {
	i.Enclosure = &Enclosure{
		URL:    url,
		Type:   enclosureType,
		Length: lengthInSeconds,
	}
}
