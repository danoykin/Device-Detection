/* *********************************************************************
 * This Source Code Form is copyright of 51Degrees Mobile Experts Limited.
 * Copyright 2017 51Degrees Mobile Experts Limited, 5 Charlotte Close,
 * Caversham, Reading, Berkshire, United Kingdom RG4 7BY
 *
 * This Source Code Form is the subject of the following patents and patent
 * applications, owned by 51Degrees Mobile Experts Limited of 5 Charlotte
 * Close, Caversham, Reading, Berkshire, United Kingdom RG4 7BY:
 * European Patent No. 2871816;
 * European Patent Application No. 17184134.9;
 * United States Patent Nos. 9,332,086 and 9,350,823; and
 * United States Patent Application No. 15/686,066.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0.
 *
 * If a copy of the MPL was not distributed with this file, You can obtain
 * one at http://mozilla.org/MPL/2.0/.
 *
 * This Source Code Form is "Incompatible With Secondary Licenses", as
 * defined by the Mozilla Public License, v. 2.0.
 ********************************************************************** */

/*
<tutorial>
Getting started example of using 51Degrees Hash Trie device detection. 
The example shows how to:
<ol>
<li>Instantiate the 51Degrees device detection provider.
<p><pre class="prettyprint lang-go">
provider := FiftyOneDegreesPatternV3.NewProvider(dataFile)
</pre></p>
<li>Produce a match for a single HTTP User-Agent header
<p><pre class="prettyprint lang-go">
match := provider.GetMatch(userAgent)
</pre></p>
<li>Extract the value of the IsMobile property
<p><pre class="prettyprint lang-go">
match.GetValue("IsMobile")
</pre></p>
</ol>
This example assumes you have the 51Degrees Go Hash Trie API installed 
correctly.
</tutorial>
*/

// Snippet Start
package main

import (
    "fmt"
    "./src/trie"
)

// Location of data file.
var dataFile = "../data/51Degrees-LiteV3.4.trie"

func main() {

    // User-Agent string of an iPhone mobile device.
    mobileUserAgent := "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) 'Version/7.0 Mobile/11D167 Safari/9537.53";

    // User-Agent string of Firefox Web browser version 41 on desktop.
    desktopUserAgent := "Mozilla/5.0 (Windows NT 6.3; WOW64; rv:41.0) Gecko/20100101 Firefox/41.0";

    // User-Agent string of a MediaHub device.
    mediaHubUserAgent := "Mozilla/5.0 (Linux; Android 4.4.2; X7 Quad Core Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Safari/537.36";

    // Provides access to device detection functions.
    provider := FiftyOneDegreesTrieV3.NewProvider(dataFile)

    fmt.Println("Starting Getting Started Example.")

    // Carries out a match for a mobile User-Agent.
    fmt.Println("\nMobile User-Agent: ", mobileUserAgent)
    match := provider.GetMatch(mobileUserAgent)
    fmt.Println("   IsMobile: ", match.GetValue("IsMobile"))
    FiftyOneDegreesTrieV3.DeleteMatch(match)

    // Carries out a macth for a desktop User-Agent.
    fmt.Println("\nDesktop User-Agent: ", desktopUserAgent)
    match = provider.GetMatch(desktopUserAgent)
    fmt.Println("   IsMobile: ", match.GetValue("IsMobile"))
    FiftyOneDegreesTrieV3.DeleteMatch(match)

    // Carries out a match for a MediaHub User-Agent.
    fmt.Println("\nMediaHub User-Agent: ", mediaHubUserAgent)
    match = provider.GetMatch(mediaHubUserAgent)
    fmt.Println("   IsMobile: ", match.GetValue("IsMobile"))
    FiftyOneDegreesTrieV3.DeleteMatch(match)
}
// Snippet End