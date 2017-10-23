package main

import (
	"testing"
)

func TestRenderBadge(t *testing.T) {
	badge := RenderBadge([]string{"test-report.xml"})
	expected := `<svg xmlns="http://www.w3.org/2000/svg" width="96" height="20">
    <title>90%</title>
    <desc>Generated with covbadger (https://github.com/imsky/covbadger)</desc>
    <linearGradient id="smooth" x2="0" y2="100%">
        <stop offset="0" stop-color="#bbb" stop-opacity=".1" />
        <stop offset="1" stop-opacity=".1" />
    </linearGradient>
    <rect rx="3" width="96" height="20" fill="#555" />
    <rect rx="3" x="60" width="36" height="20" fill="#6ccb08" />
    <rect x="60" width="4" height="20" fill="#6ccb08" />
    <rect rx="3" width="96" height="20" fill="url(#smooth)" />
    <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,sans-serif" font-size="11">
        <text x="30" y="15" fill="#010101" fill-opacity=".3">coverage</text>
        <text x="30" y="14">coverage</text>
        <text x="78" y="15" fill="#010101" fill-opacity=".3">90%</text>
        <text x="78" y="14">90%</text>
    </g>
</svg>`

	if badge != expected {
		t.Errorf("RenderBadge output is incorrect")
	}
}
