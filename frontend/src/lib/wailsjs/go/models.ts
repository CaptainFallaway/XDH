export namespace internal {
	
	export class Session {
	    uid: string;
	    reading: string;
	    date: number;
	    location: string;
	    instrumentSerial: string;
	
	    static createFrom(source: any = {}) {
	        return new Session(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.uid = source["uid"];
	        this.reading = source["reading"];
	        this.date = source["date"];
	        this.location = source["location"];
	        this.instrumentSerial = source["instrumentSerial"];
	    }
	}

}

