/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
* The `forbiddenmarkers` linter ensures that types and fields do not contain any markers
* that are forbidden.
*
* By default, `forbiddenmarkers` is not enabled.
*
* It can be configured with a list of marker identifiers and optionally their attributes and values that are forbidden
* ```yaml
* lintersConfig:
*   forbiddenMarkers:
*     markers:
*       - identifier: forbidden:marker
*       - identifier: forbidden:withAttribute
*         attributes:
*           - name: fruit
*       - identifier: forbidden:withMultipleAttributes
*         attributes:
*           - name: fruit
*           - name: color
*       - identifier: forbidden:withAttributeValues
*         attributes:
*           - name: fruit
*             values:
*               - apple
*               - banana
*               - orange
*       - identifier: forbidden:withMultipleAttributesValues
*         attributes:
*           - name: fruit
*             values:
*               - apple
*               - banana
*               - orange
*           - name: color
*             values:
*               - red
*               - green
*               - blue
* ```
*
* Using the config above, the following examples would be forbidden:
* - `+forbidden:marker` (all instances, including if they have attributes and values)
* - `+forbidden:withAttribute:fruit=*,*` (all instances of this marker containing the attribute 'fruit')
* - `+forbidden:withMultipleAttributes:fruit=*,color=*,*` (all instances of this marker containing both the 'fruit' AND 'color' attributes)
* - `+forbidden:withAttributeValues:fruit={apple || banana || orange},*` (all instances of this marker containing the 'fruit' attribute where the value is set to one of 'apple', 'banana', or 'orange')
* - `+forbidden:withMultipleAttributesValues:fruit={apple || banana || orange},color={red || green || blue},*` (all instances of this marker containing the 'fruit' attribute where the value is set to one of 'apple', 'banana', or 'orange' AND the 'color' attribute where the value is set to one of 'red', 'green', or 'blue')
*
* Fixes are suggested to remove all markers that are forbidden.
*
 */
package forbiddenmarkers
