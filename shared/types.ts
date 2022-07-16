/* Do not change, this code is generated from Golang structs */


export class Feed {
    url: string;
    title: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.url = source["url"];
        this.title = source["title"];
    }
}
export class Event {
    user_id: string;
    feeds: Feed[];

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.user_id = source["user_id"];
        this.feeds = this.convertValues(source["feeds"], Feed);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
export class PostContent {
    title: string;
    byline: string;
    site: string;
    text: string;
    favicon: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.title = source["title"];
        this.byline = source["byline"];
        this.site = source["site"];
        this.text = source["text"];
        this.favicon = source["favicon"];
    }
}