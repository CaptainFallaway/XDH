export namespace internal {
	
	export class Scan {
	    reading: number;
	    duration: number;
	    operator: string;
	    date: number;
	    pb: number;
	    zn: number;
	    cu: number;
	    sn: number;
	
	    static createFrom(source: any = {}) {
	        return new Scan(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.reading = source["reading"];
	        this.duration = source["duration"];
	        this.operator = source["operator"];
	        this.date = source["date"];
	        this.pb = source["pb"];
	        this.zn = source["zn"];
	        this.cu = source["cu"];
	        this.sn = source["sn"];
	    }
	}
	export class Grouping {
	    index: number;
	    boatID: string;
	    firstDate: number;
	    lastDate: number;
	    unit: string;
	    scans: Scan[];
	    invalidScans: Scan[];
	    errorNotes: string[];
	    violations: {[key: string]: number};
	    operators: string[];
	
	    static createFrom(source: any = {}) {
	        return new Grouping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.boatID = source["boatID"];
	        this.firstDate = source["firstDate"];
	        this.lastDate = source["lastDate"];
	        this.unit = source["unit"];
	        this.scans = this.convertValues(source["scans"], Scan);
	        this.invalidScans = this.convertValues(source["invalidScans"], Scan);
	        this.errorNotes = source["errorNotes"];
	        this.violations = source["violations"];
	        this.operators = source["operators"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
	
	export class SessionInfo {
	    uid: string;
	    surveyor: string;
	    date: number;
	    location: string;
	    instrumentSerial: string;
	
	    static createFrom(source: any = {}) {
	        return new SessionInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uid = source["uid"];
	        this.surveyor = source["surveyor"];
	        this.date = source["date"];
	        this.location = source["location"];
	        this.instrumentSerial = source["instrumentSerial"];
	    }
	}

}

