package models

type Image struct {
	ID       string
	URL      string
	Alt      string
	Location string
}

var images = []Image{
	{URL: "/assets/images/tower1.jpg", Alt: "Bromölla", Location: "Bromölla", ID: "bromölla"},
	{URL: "/assets/images/tower2.jpg", Alt: "Kristanstad", Location: "Kristanstad", ID: "kristanstad"},
	{URL: "/assets/images/tower3.jpg", Alt: "Malmö - Hyllie", Location: "Malmö - Hyllie", ID: "malmöhyllie"},
	{URL: "/assets/images/tower4.jpg", Alt: "Malmö - Södervärn", Location: "Malmö - Södervärn", ID: "malmösödervärn"},
	{URL: "/assets/images/tower5.jpg", Alt: "Oxie", Location: "Oxie", ID: "oxie"},
	{URL: "/assets/images/tower6.jpg", Alt: "Ystad", Location: "Ystad", ID: "ystad"},
	{URL: "/assets/images/tower7.jpg", Alt: "Hälsingborg", Location: "Hälsingborg", ID: "hälsingborg"},
}
