package gonvim

import (
	"fmt"
	"sync"
)

var svgsOnce sync.Once
var svgs map[string]*SvgXML

// SvgXML is
type SvgXML struct {
	xml       string
	width     int
	height    int
	thickness float64
	color     *RGBA
}

// Svg is
type Svg struct {
	width  int
	height int
	color  *RGBA
	bg     *RGBA
	name   string
}

func getSvgs() map[string]*SvgXML {
	svgsOnce.Do(func() {
		initSVGS()
	})
	return svgs
}

func getSvg(name string, color *RGBA) string {
	if editor == nil {
		return ""
	}
	svgsOnce.Do(func() {
		initSVGS()
	})
	svg := svgs[name]
	if svg == nil {
		svg = svgs["default"]
	}
	if color == nil {
		color = svg.color
	}
	if color == nil {
		color = editor.Foreground
	}

	return fmt.Sprintf(svg.xml, color.Hex())
}

func initSVGS() {
	blue := newRGBA(81, 154, 186, 1)
	svgs = map[string]*SvgXML{}
	svgs["fire"] = &SvgXML{
		width:  1792,
		height: 1792,
		xml: `<?xml version="1.0" encoding="utf-8"?>
<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1600 1696v64q0 13-9.5 22.5t-22.5 9.5h-1344q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h1344q13 0 22.5 9.5t9.5 22.5zm-256-1056q0 78-24.5 144t-64 112.5-87.5 88-96 77.5-87.5 72-64 81.5-24.5 96.5q0 96 67 224l-4-1 1 1q-90-41-160-83t-138.5-100-113.5-122.5-72.5-150.5-27.5-184q0-78 24.5-144t64-112.5 87.5-88 96-77.5 87.5-72 64-81.5 24.5-96.5q0-94-66-224l3 1-1-1q90 41 160 83t138.5 100 113.5 122.5 72.5 150.5 27.5 184z"/></svg>`,
	}
	svgs["comment"] = &SvgXML{
		width:  1792,
		height: 1792,
		color:  blue,
		xml: `<?xml version="1.0" encoding="utf-8"?>
<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1792 896q0 174-120 321.5t-326 233-450 85.5q-70 0-145-8-198 175-460 242-49 14-114 22-17 2-30.5-9t-17.5-29v-1q-3-4-.5-12t2-10 4.5-9.5l6-9 7-8.5 8-9q7-8 31-34.5t34.5-38 31-39.5 32.5-51 27-59 26-76q-157-89-247.5-220t-90.5-281q0-130 71-248.5t191-204.5 286-136.5 348-50.5q244 0 450 85.5t326 233 120 321.5z"/></svg>`,
	}
	svgs["cross"] = &SvgXML{
		width:     100,
		height:    100,
		thickness: 2,
		xml:       `<?xml version="1.0" encoding="utf-8"?><svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1490 1322q0 40-28 68l-136 136q-28 28-68 28t-68-28l-294-294-294 294q-28 28-68 28t-68-28l-136-136q-28-28-28-68t28-68l294-294-294-294q-28-28-28-68t28-68l136-136q28-28 68-28t68 28l294 294 294-294q28-28 68-28t68 28l136 136q28 28 28 68t-28 68l-294 294 294 294q28 28 28 68z"/></svg>`,
	}
	svgs["default"] = &SvgXML{
		width:     200,
		height:    200,
		thickness: 0.5,
		color:     editor.Foreground,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200"><g><path fill="%s" d="M22.34999879828441,17.7129974066314 h147.91225647948028 v20.748584330809138 H22.34999879828441 V17.7129974066314 zM22.34999879828441,65.18324337560382 h126.22055467908892 v20.748584330809138 H22.34999879828441 V65.18324337560382 zM22.34999879828441,113.91097930401922 h155.3000099912078 v20.748584330809138 H22.34999879828441 V113.91097930401922 zM22.34999879828441,161.538411517922 h91.01083581468554 v20.748584330809138 H22.34999879828441 V161.538411517922 z"/></g></svg>`,
	}
	svgs["empty"] = &SvgXML{
		width:  200,
		height: 200,
		color:  editor.Foreground,
		xml:    `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200"><g><path fill="%s" d=""/></g></svg>`,
	}
	svgs["folder"] = &SvgXML{
		width:     1792,
		height:    1792,
		thickness: 0.5,
		color:     editor.Foreground,
		xml:       `<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><g><path fill="%s" d="M1600 1312v-704q0-40-28-68t-68-28h-704q-40 0-68-28t-28-68v-64q0-40-28-68t-68-28h-320q-40 0-68 28t-28 68v960q0 40 28 68t68 28h1216q40 0 68-28t28-68zm128-704v704q0 92-66 158t-158 66h-1216q-92 0-158-66t-66-158v-960q0-92 66-158t158-66h320q92 0 158 66t66 158v32h672q92 0 158 66t66 158z"/></g></svg>`,
	}
	svgs["git"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 0,
		color:     editor.Foreground,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M124.737 58.378l-55.116-55.114c-3.172-3.174-8.32-3.174-11.497 0l-11.444 11.446 14.518 14.518c3.375-1.139 7.243-.375 9.932 2.314 2.703 2.706 3.461 6.607 2.294 9.993l13.992 13.993c3.385-1.167 7.292-.413 9.994 2.295 3.78 3.777 3.78 9.9 0 13.679-3.78 3.78-9.901 3.78-13.683 0-2.842-2.844-3.545-7.019-2.105-10.521l-13.048-13.048-.002 34.341c.922.455 1.791 1.063 2.559 1.828 3.778 3.777 3.778 9.898 0 13.683-3.779 3.777-9.904 3.777-13.679 0-3.778-3.784-3.778-9.905 0-13.683.934-.933 2.014-1.638 3.167-2.11v-34.659c-1.153-.472-2.231-1.172-3.167-2.111-2.862-2.86-3.551-7.06-2.083-10.576l-14.313-14.313-37.792 37.79c-3.175 3.177-3.175 8.325 0 11.5l55.117 55.114c3.174 3.174 8.32 3.174 11.499 0l54.858-54.858c3.174-3.176 3.174-8.327-.001-11.501z"/></g></svg>`,
	}
	svgs["check"] = &SvgXML{
		width:     1792,
		height:    1792,
		thickness: 0,
		color:     newRGBA(141, 193, 73, 1),
		xml:       `<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><g><path fill="%s" d="M1671 566q0 40-28 68l-724 724-136 136q-28 28-68 28t-68-28l-136-136-362-362q-28-28-28-68t28-68l136-136q28-28 68-28t68 28l294 295 656-657q28-28 68-28t68 28l136 136q28 28 28 68z"/></g></svg>`,
	}
	svgs["exclamation"] = &SvgXML{
		width:     1792,
		height:    1792,
		thickness: 0,
		color:     newRGBA(203, 203, 65, 1),
		xml:       `<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><g><path fill="%s" d="M1088 1248v224q0 26-19 45t-45 19h-256q-26 0-45-19t-19-45v-224q0-26 19-45t45-19h256q26 0 45 19t19 45zm30-1056l-28 768q-1 26-20.5 45t-45.5 19h-256q-26 0-45.5-19t-20.5-45l-28-768q-1-26 17.5-45t44.5-19h320q26 0 44.5 19t17.5 45z"/></g></svg>`,
	}
	svgs["sh"] = &SvgXML{
		width:     18,
		height:    18,
		thickness: 0,
		color:     editor.Foreground,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 18 18"><g><path fill="%s" d="M1.9426860809326185,2.020794440799331 V15.979206420212122 H16.057314060551434 V2.020794440799331 zm2.1962214082875993,1.8041484023191938 l1.488652194536292,1.1762194099782244 l-1.488652194536292,1.2543270180082322 zm0,3.606765445891675 H13.234695018669578 v1.0200022335531784 H4.138907489220218 zm0,2.9007278261308826 h5.723346475939733 v1.0200022335531784 H4.138907489220218 zm0,2.824150479043169 h7.60560350567875 v1.018470796391866 H4.138907489220218 z" /></g></svg>`,
	}
	svgs["py"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 0,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M49.33 62h29.159c8.117 0 14.511-6.868 14.511-15.019v-27.798c0-7.912-6.632-13.856-14.555-15.176-5.014-.835-10.195-1.215-15.187-1.191-4.99.023-9.612.448-13.805 1.191-12.355 2.181-14.453 6.751-14.453 15.176v10.817h29v4h-40.224000000000004c-8.484 0-15.914 5.108-18.237 14.811-2.681 11.12-2.8 17.919 0 29.53 2.075 8.642 7.03 14.659 15.515 14.659h9.946v-13.048c0-9.637 8.428-17.952 18.33-17.952zm-1.838-39.11c-3.026 0-5.478-2.479-5.478-5.545 0-3.079 2.451-5.581 5.478-5.581 3.015 0 5.479 2.502 5.479 5.581-.001 3.066-2.465 5.545-5.479 5.545zM122.281 48.811c-2.098-8.448-6.103-14.811-14.599-14.811h-10.682v12.981c0 10.05-8.794 18.019-18.511 18.019h-29.159c-7.988 0-14.33 7.326-14.33 15.326v27.8c0 7.91 6.745 12.564 14.462 14.834 9.242 2.717 17.994 3.208 29.051 0 7.349-2.129 14.487-6.411 14.487-14.834v-11.126h-29v-4h43.682c8.484 0 11.647-5.776 14.599-14.66 3.047-9.145 2.916-17.799 0-29.529zm-41.955 55.606c3.027 0 5.479 2.479 5.479 5.547 0 3.076-2.451 5.579-5.479 5.579-3.015 0-5.478-2.502-5.478-5.579 0-3.068 2.463-5.547 5.478-5.547z"/></g></svg>`,
	}
	svgs["pyc"] = svgs["py"]
	svgs["c"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 1,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M97.5221405461895,45.03725476832212 L108.76640973908529,40.28547464292772 C108.56169322951193,38.08885210339795 94.3709708336648,16.355134752322773 69.08566426354695,16.60775090230729 C43.80035769342909,16.86037464885613 18.857250847109363,37.373658362115926 19.237894612168745,64.2531288574854 C19.61853837722812,91.1325841597262 44.09328059332486,111.39442159390717 68.83481490446007,111.39442919047148 C93.57635717502413,111.39443678703581 108.34661354359856,92.47813321193519 107.91159096075516,92.0739200243284 C107.47656837791175,91.66970683672162 100.93955315316323,85.45629455970422 100.26450603444523,85.0460952794112 C99.58945891572723,84.63589599911819 90.07591976589798,98.4070389619417 68.86888921929224,98.15440761882854 C47.66185867268652,97.90177627571538 33.5390540830676,80.57280520208417 33.375925589144344,64.50575260403423 C33.212797095221084,48.43870000598427 47.89705183525975,31.110936786080387 69.10406646300783,30.858313039531545 C90.31109700961355,30.605681696418387 93.33611773457105,42.36367434123211 97.5221405461895,45.03725476832212 z"></path></g></svg>`,
	}
	svgs["cpp"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 0.5,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M93.76617750047221,41.61478427062541 L107.12808612823277,35.96809876661383 C106.88481524520896,33.35778508954351 90.02154625347892,7.530944731601865 59.97424391494265,7.831136220802641 C29.926941576406406,8.131336737232896 0.28628562487806786,32.507898515483696 0.7386162406964498,64.4495945528423 C1.1909468565148238,96.39127253574195 30.275030808017725,120.46903301813222 59.67615195734626,120.4690420453617 C89.07728256510728,120.4690510725912 106.62922948311446,97.99022958500447 106.11227885668883,97.5098907042571 C105.59532823026318,97.02955182350973 97.82719331345099,89.64596417651826 97.02501419758264,89.1585118389389 C96.22283508171427,88.67105950135958 84.91760584307633,105.03572943342894 59.71664350673292,104.73551988976921 C34.51568117038954,104.43531034610947 17.733120982891318,83.84276456774445 17.539270409502,64.74979506927255 C17.34541983611268,45.656825570800656 34.79516839132544,25.065715121923326 59.996111810803896,24.76551460549306 C85.19707414714729,24.465305061833327 88.79179869798797,38.437686963286254 93.76617750047221,41.61478427062541 zM102.17347564698959,64.32615415124337 L127.35898957621174,63.97142860294444 zM69.18400149080784,64.32615755108435 L94.36951542002998,63.971432002785434 zM69.5387267330024,64.32615641780403 L94.72424066222455,63.9714308695051 zM102.52820088918415,64.32615301796304 L127.71371481840629,63.97142746966412 zM94.45454502105713,60.50000059604645 l30,0 l0,7 l-30,0 zM105.95454502105713,49.00000059604645 l7,0 l0,30 l-7,0 zM58.69696617126465,60.499997079372406 l30,0 l0,7 l-30,0 zM70.19696617126465,48.999997079372406 l7,0 l0,30 l-7,0 z"></path></g></svg>`,
	}
	svgs["go"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 0,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M108.2 64.8c-.1-.1-.2-.2-.4-.2l-.1-.1c-.1-.1-.2-.1-.2-.2l-.1-.1c-.1 0-.2-.1-.2-.1l-.2-.1c-.1 0-.2-.1-.2-.1l-.2-.1c-.1 0-.2-.1-.2-.1-.1 0-.1 0-.2-.1l-.3-.1c-.1 0-.1 0-.2-.1l-.3-.1h-.1l-.4-.1h-.2c-.1 0-.2 0-.3-.1h-2.3c-.6-13.3.6-26.8-2.8-39.6 12.9-4.6 2.8-22.3-8.4-14.4-7.4-6.4-17.6-7.8-28.3-7.8-10.5.7-20.4 2.9-27.4 8.4-2.8-1.4-5.5-1.8-7.9-1.1v.1c-.1 0-.3.1-.4.2-.1 0-.3.1-.4.2h-.1c-.1 0-.2.1-.4.2h-.1l-.3.2h-.1l-.3.2h-.1l-.3.2s-.1 0-.1.1l-.3.2s-.1 0-.1.1l-.3.2s-.1 0-.1.1l-.3.2-.1.1c-.1.1-.2.1-.2.2l-.1.1-.2.2-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1-.1.3s0 .1-.1.1l-.1.3s0 .1-.1.1l-.1.3s0 .1-.1.1l-.1.3s0 .1-.1.1c.4.3.4.4.4.4v.1l-.1.3v.1c0 .1 0 .2-.1.3v3.1c0 .1 0 .2.1.3v.1l.1.3v.1l.1.3s0 .1.1.1l.1.3s0 .1.1.1l.1.3s0 .1.1.1l.2.3s0 .1.1.1l.2.3s0 .1.1.1l.2.3.1.1.3.3.3.3h.1c1 .9 2 1.6 4 2.2v-.2c-4.2 12.6-.7 25.3-.5 38.3-.6 0-.7.4-1.7.5h-.5c-.1 0-.3 0-.5.1-.1 0-.3 0-.4.1l-.4.1h-.1l-.4.1h-.1l-.3.1h-.1l-.3.1s-.1 0-.1.1l-.3.1-.2.1c-.1 0-.2.1-.2.1l-.2.1-.2.1c-.1 0-.2.1-.2.1l-.2.1-.4.3c-.1.1-.2.2-.3.2l-.4.4-.1.1c-.1.2-.3.4-.4.5l-.2.3-.3.6-.1.3v.3c0 .5.2.9.9 1.2.2 3.7 3.9 2 5.6.8l.1-.1c.2-.2.5-.3.6-.3h.1l.2-.1c.1 0 .1 0 .2-.1.2-.1.4-.1.5-.2.1 0 .1-.1.1-.2l.1-.1c.1-.2.2-.6.2-1.2l.1-1.3v1.8c-.5 13.1-4 30.7 3.3 42.5 1.3 2.1 2.9 3.9 4.7 5.4h-.5c-.2.2-.5.4-.8.6l-.9.6-.3.2-.6.4-.9.7-1.1 1c-.2.2-.3.4-.4.5l-.4.6-.2.3c-.1.2-.2.4-.2.6l-.1.3c-.2.8 0 1.7.6 2.7l.4.4h.2c.1 0 .2 0 .4.1.2.4 1.2 2.5 3.9.9 2.8-1.5 4.7-4.6 8.1-5.1l-.5-.6c5.9 2.8 12.8 4 19 4.2 8.7.3 18.6-.9 26.5-5.2 2.2.7 3.9 3.9 5.8 5.4l.1.1.1.1.1.1.1.1s.1 0 .1.1c0 0 .1 0 .1.1 0 0 .1 0 .1.1h2.1000000000000005s.1 0 .1-.1h.1s.1 0 .1-.1h.1s.1 0 .1-.1c0 0 .1 0 .1-.1l.1-.1s.1 0 .1-.1l.1-.1h.1l.2-.2.2-.1h.1l.1-.1h.1l.1-.1.1-.1.1-.1.1-.1.1-.1.1-.1.1-.1v-.1s0-.1.1-.1v-.1s0-.1.1-.1v-.1s0-.1.1-.1v-1.4000000000000001s-.3 0-.3-.1l-.3-.1v-.1l.3-.1s.2 0 .2-.1l.1-.1v-2.1000000000000005s0-.1-.1-.1v-.1s0-.1-.1-.1v-.1s0-.1-.1-.1c0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1l-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1v-.1l-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1c2-1.9 3.8-4.2 5.1-6.9 5.9-11.8 4.9-26.2 4.1-39.2h.1c.1 0 .2.1.2.1h.30000000000000004s.1 0 .1.1h.1s.1 0 .1.1l.2.1c1.7 1.2 5.4 2.9 5.6-.8 1.6.6-.3-1.8-1.3-2.5zm-72.2-41.8c-3.2-16 22.4-19 23.3-3.4.8 13-20 16.3-23.3 3.4zm36.1 15c-1.3 1.4-2.7 1.2-4.1.7 0 1.9.4 3.9.1 5.9-.5.9-1.5 1-2.3 1.4-1.2-.2-2.1-.9-2.6-2l-.2-.1c-3.9 5.2-6.3-1.1-5.2-5-1.2.1-2.2-.2-3-1.5-1.4-2.6.7-5.8 3.4-6.3.7 3 8.7 2.6 10.1-.2 3.1 1.5 6.5 4.3 3.8 7.1zm-7-17.5c-.9-13.8 20.3-17.5 23.4-4 3.5 15-20.8 18.9-23.4 4zM41.7 17c-1.9 0-3.5 1.7-3.5 3.8 0 2.1 1.6 3.8 3.5 3.8s3.5-1.7 3.5-3.8c0-2.1-1.5-3.8-3.5-3.8zm1.6 5.7c-.5 0-.8-.4-.8-1 0-.5.4-1 .8-1 .5 0 .8.4.8 1 0 .5-.3 1-.8 1zM71.1 16.1c-1.9 0-3.4 1.7-3.4 3.8 0 2.1 1.5 3.8 3.4 3.8s3.4-1.7 3.4-3.8c0-2.1-1.5-3.8-3.4-3.8zm1.6 5.6c-.4 0-.8-.4-.8-1 0-.5.4-1 .8-1s.8.4.8 1-.4 1-.8 1z"/></g></svg>`,
	}
}
