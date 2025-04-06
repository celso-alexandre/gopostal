package zipcode

func NormalizeBrazilZipCode(zipcode string) string {
	if zipcode == "" {
		return ""
	}
	normalized := ""
	for _, char := range zipcode {
		if char >= '0' && char <= '9' {
			normalized += string(char)
		}
	}
	if len(normalized) < 8 {
		return normalized
	}
	return normalized[:5] + "-" + normalized[5:]
}
