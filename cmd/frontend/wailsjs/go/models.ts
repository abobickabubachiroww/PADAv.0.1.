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
	    GroupName: string;
	    Dadt: string;
	    CnStr: number;
	    RVP: number;
	    TPI: number;
	    RVPD: number;
	    TPID: number;
	    GDP: number;
	    PSS: number;
	    PSA: number;
	    IIS: number;
	    IPE: number;
	    L0: number;
	    L1: number;
	    L2: number;
	    L3: number;
	    Special: number;
	
	    static createFrom(source: any = {}) {
	        return new PersonalMetrics(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.FIO = source["FIO"];
	        this.GroupName = source["GroupName"];
	        this.Dadt = source["Dadt"];
	        this.CnStr = source["CnStr"];
	        this.RVP = source["RVP"];
	        this.TPI = source["TPI"];
	        this.RVPD = source["RVPD"];
	        this.TPID = source["TPID"];
	        this.GDP = source["GDP"];
	        this.PSS = source["PSS"];
	        this.PSA = source["PSA"];
	        this.IIS = source["IIS"];
	        this.IPE = source["IPE"];
	        this.L0 = source["L0"];
	        this.L1 = source["L1"];
	        this.L2 = source["L2"];
	        this.L3 = source["L3"];
	        this.Special = source["Special"];
	    }
	}

}

export namespace ui {
	
	export class GroupReportFilter {
	    institute: string;
	    year: number;
	    half: string;
	
	    static createFrom(source: any = {}) {
	        return new GroupReportFilter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.institute = source["institute"];
	        this.year = source["year"];
	        this.half = source["half"];
	    }
	}

}

