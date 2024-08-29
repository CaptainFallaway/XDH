export namespace data_pipeline {
	
	export class MetalValue {
	    value: number;
	    isLod: boolean;
	
	    static createFrom(source: any = {}) {
	        return new MetalValue(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.value = source["value"];
	        this.isLod = source["isLod"];
	    }
	}
	export class ScanRow {
	    index: number;
	    reading: number;
	    time: string;
	    type: string;
	    duration: number;
	    unit: string;
	    sigmaValue: number;
	    sequence: string;
	    user1: string;
	    flags: string;
	    boat: string;
	    operator: string;
	    userLogin: string;
	    pb: MetalValue;
	    pbError: number;
	    zn: MetalValue;
	    znError: number;
	    cu: MetalValue;
	    cuError: number;
	    sn: MetalValue;
	    snError: number;
	
	    static createFrom(source: any = {}) {
	        return new ScanRow(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.reading = source["reading"];
	        this.time = source["time"];
	        this.type = source["type"];
	        this.duration = source["duration"];
	        this.unit = source["unit"];
	        this.sigmaValue = source["sigmaValue"];
	        this.sequence = source["sequence"];
	        this.user1 = source["user1"];
	        this.flags = source["flags"];
	        this.boat = source["boat"];
	        this.operator = source["operator"];
	        this.userLogin = source["userLogin"];
	        this.pb = this.convertValues(source["pb"], MetalValue);
	        this.pbError = source["pbError"];
	        this.zn = this.convertValues(source["zn"], MetalValue);
	        this.znError = source["znError"];
	        this.cu = this.convertValues(source["cu"], MetalValue);
	        this.cuError = source["cuError"];
	        this.sn = this.convertValues(source["sn"], MetalValue);
	        this.snError = source["snError"];
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
	export class Grouping {
	    index: number;
	    firstDate: string;
	    lastDate: string;
	    unit: string;
	    scans: ScanRow[];
	    invalidScans: ScanRow[];
	    errorNotes: string[];
	    violations: {[key: string]: number};
	    boatID: string;
	    operators: string[];
	
	    static createFrom(source: any = {}) {
	        return new Grouping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.index = source["index"];
	        this.firstDate = source["firstDate"];
	        this.lastDate = source["lastDate"];
	        this.unit = source["unit"];
	        this.scans = this.convertValues(source["scans"], ScanRow);
	        this.invalidScans = this.convertValues(source["invalidScans"], ScanRow);
	        this.errorNotes = source["errorNotes"];
	        this.violations = source["violations"];
	        this.boatID = source["boatID"];
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
	

}

