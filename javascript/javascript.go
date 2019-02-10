package javascript

import "fmt"

func WindowSize() string {
	return fmt.Sprintf(`(function () {
    return {
        height: window.innerHeight,
        width:  window.innerWidth,
    };
})();`)
}

// source: https://stackoverflow.com/questions/8813051/determine-which-element-the-mouse-pointer-is-on-top-of-in-javascript
func GetElementsXPathForPoint(x, y int64) string {
	return fmt.Sprintf(`(function (x, y) {

    let getPathTo = function(element) {
        if (element.id!=='')
            return 'id("'+element.id+'")';
        if (element===document.body)
            return element.tagName;

        var ix= 0;
        var siblings= element.parentNode.childNodes;
        for (var i= 0; i<siblings.length; i++) {
            var sibling= siblings[i];
            if (sibling===element)
                return getPathTo(element.parentNode)+'/'+element.tagName+'['+(ix+1)+']';
            if (sibling.nodeType===1 && sibling.tagName===element.tagName)
                ix++;
        }
    }

	console.log("chromedpcv: "+x+" : "+y );
    var element, elements = [];
    var old_visibility = [];
    while (true) {
        element = document.elementFromPoint(x, y);
        if (!element || element === document.documentElement) {
            break;
        }
        elements.push(element);
        old_visibility.push(element.style.visibility);
        element.style.visibility = 'hidden'; // Temporarily hide the element (without changing the layout)
    }
    for (var k = 0; k < elements.length; k++) {
        elements[k].style.visibility = old_visibility[k];
    }
    elements.reverse();

	var xpaths = [];
	for (var k=0; k < elements.length; k++ ) {
		xpaths.push(getPathTo(elements[k]));
	}
    return xpaths.reverse();
	})(%d, %d);`, x, y)
}

// source: https://stackoverflow.com/questions/48903805/how-to-connect-onmousemove-with-onmousedown
const LogMouseClicksInConsole = `(function() {
    document.onmouseclick = handleMouseMove;
    function handleMouseMove(event) {
        var eventDoc, doc, body;
        event = event || browserWindow.event;
        if (event.pageX == null && event.clientX != null) {
            eventDoc = (event.target && event.target.ownerDocument) || document;
            doc = eventDoc.documentElement;
            body = eventDoc.body;
            event.pageX = event.clientX +
              (doc && doc.scrollLeft || body && body.scrollLeft || 0) -
              (doc && doc.clientLeft || body && body.clientLeft || 0);
            event.pageY = event.clientY +
              (doc && doc.scrollTop  || body && body.scrollTop  || 0) -
              (doc && doc.clientTop  || body && body.clientTop  || 0 );
        }
		console.log( event.pageX + " : " + event.pageY + " chromedpcv:");
        // Use event.pageX / event.pageY here
    }
})();`
