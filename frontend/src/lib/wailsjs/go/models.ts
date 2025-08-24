export namespace api {
	
	export class Grouping {
	    validScans: models.Scan[];
	    invalidScans: models.Scan[];
	
	    static createFrom(source: any = {}) {
	        return new Grouping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.validScans = this.convertValues(source["validScans"], models.Scan);
	        this.invalidScans = this.convertValues(source["invalidScans"], models.Scan);
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
	export class Survey {
	    surveyor: string;
	    date: number;
	    location: string;
	    westCoastFlag: boolean;
	    instrumentSerial: string;
	
	    static createFrom(source: any = {}) {
	        return new Survey(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.surveyor = source["surveyor"];
	        this.date = source["date"];
	        this.location = source["location"];
	        this.westCoastFlag = source["westCoastFlag"];
	        this.instrumentSerial = source["instrumentSerial"];
	    }
	}

}

export namespace models {
	
	export class Scan {
	    reading: number;
	    duration: number;
	    operator: string;
	    date: number;
	    pb: number;
	    zn: number;
	    cu: number;
	    sn: number;
	    violations: Record<string, boolean>;
	
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
	        this.violations = source["violations"];
	    }
	}
	export class Grouping {
	    uid: string;
	    index: number;
	    boatID: string;
	    firstDate: number;
	    lastDate: number;
	    unit: string;
	    errorNotes: string[];
	    violations: Record<string, number>;
	    validScans: Scan[];
	    invalidScans: Scan[];
	    operators: string[];
	
	    static createFrom(source: any = {}) {
	        return new Grouping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uid = source["uid"];
	        this.index = source["index"];
	        this.boatID = source["boatID"];
	        this.firstDate = source["firstDate"];
	        this.lastDate = source["lastDate"];
	        this.unit = source["unit"];
	        this.errorNotes = source["errorNotes"];
	        this.violations = source["violations"];
	        this.validScans = this.convertValues(source["validScans"], Scan);
	        this.invalidScans = this.convertValues(source["invalidScans"], Scan);
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
	
	export class Survey {
	    uid: string;
	    surveyor: string;
	    date: number;
	    location: string;
	    instrumentSerial: string;
	    westCoastFlag: boolean;
	    groupingIds: string[];
	
	    static createFrom(source: any = {}) {
	        return new Survey(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uid = source["uid"];
	        this.surveyor = source["surveyor"];
	        this.date = source["date"];
	        this.location = source["location"];
	        this.instrumentSerial = source["instrumentSerial"];
	        this.westCoastFlag = source["westCoastFlag"];
	        this.groupingIds = source["groupingIds"];
	    }
	}

}

