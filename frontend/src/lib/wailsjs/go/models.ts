export namespace data_pipeline {
	
	export class Grouping {
	
	
	    static createFrom(source: any = {}) {
	        return new Grouping(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

