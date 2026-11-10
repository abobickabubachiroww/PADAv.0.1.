export namespace model {
	
	export class GroupReport {
	    ID: number;
	    Institute: string;
	    Group_name: string;
	    Starosta: string;
	    Students_list: number;
	    Students_fact: number;
	    Missed_dln: number;
	    Quality_pct: number;
	    Measured_at: string;
	    Starosta_was: string;
	
	    static createFrom(source: any = {}) {
	        return new GroupReport(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Institute = source["Institute"];
	        this.Group_name = source["Group_name"];
	        this.Starosta = source["Starosta"];
	        this.Students_list = source["Students_list"];
	        this.Students_fact = source["Students_fact"];
	        this.Missed_dln = source["Missed_dln"];
	        this.Quality_pct = source["Quality_pct"];
	        this.Measured_at = source["Measured_at"];
	        this.Starosta_was = source["Starosta_was"];
	    }
	}
	export class PersonalMetrics {
	    ID: number;
	    FIO: string;
	    Dadt: string;
	    Smar: number;
	    Smap: number;
	    Smsd: number;
	    PctSmar: number;
	    PctSmap: number;
	    Rvp: number;
	    Tpi: number;
	    Rvpd: number;
	    Tpid: number;
	    Gdp: number;
	    Pss: number;
	    Psa: number;
	    Iis: number;
	    Ipe: number;
	    L04: number;
	    L12: number;
	
	    static createFrom(source: any = {}) {
	        return new PersonalMetrics(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.FIO = source["FIO"];
	        this.Dadt = source["Dadt"];
	        this.Smar = source["Smar"];
	        this.Smap = source["Smap"];
	        this.Smsd = source["Smsd"];
	        this.PctSmar = source["PctSmar"];
	        this.PctSmap = source["PctSmap"];
	        this.Rvp = source["Rvp"];
	        this.Tpi = source["Tpi"];
	        this.Rvpd = source["Rvpd"];
	        this.Tpid = source["Tpid"];
	        this.Gdp = source["Gdp"];
	        this.Pss = source["Pss"];
	        this.Psa = source["Psa"];
	        this.Iis = source["Iis"];
	        this.Ipe = source["Ipe"];
	        this.L04 = source["L04"];
	        this.L12 = source["L12"];
	    }
	}

}

