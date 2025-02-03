name Donald Galbraith
gender male
tags 'Argyll:Killean and Kilchenzie' Argyll:Campbeltown
external {
    familysearch LZ68-L2C
}

baptism -date 8-Aug-1819 -location 'killean and kilchenzie' -ref donald1819-baptism
death -date 20-nov-1897 -location campbeltown -ref death
note {
    Deaths, marriages, and decendants not investigated. 
}
body {
Census of 1881: 6, High Street, Campbeltown, Argyllshire, Scotland:$ref[census1881]
$csvtable{
First,Last,Role,Status,Age,Year,Place,Job
Donald,Galbraith,Head,Married,60,1821,BALACHANTYNE K PARISH GAELIC,Cartwright
Isabella,Galbraith,Wife,Married,63,1818,CAMPBELTOWN GAELIC,Cartwrights wife
Donald,Galbraith,Son,Unmarried,23,1858,CAMPBELTOWN GAELIC,Laborer
Alexander,Galbraith,Son,Unmarried,21,1860,CAMPBELTOWN GAELIC,Mason
}
}
partner {
    name Isabella McLean
    gender female
    birth -date 10-mar-1816 -location campbeltown  -note 'daughter of Donald McLean and Agnes Stewart'
    marriage -date 13-dec-1847 -location campbeltown -ref mclean-marriage
    death -date 30-dec-1887 -location campbeltown -ref mclean-death
    child {
        name Agnes Galbreath
        gender male
        baptism -date 25-jan-1848 -location campbeltown -ref agnes1848-baptism
    }
    child {
        name Isabella Galbraith
        gender female
        baptism -date 22-aug-1850 -location campbeltown -ref isabella1850-baptism
        death -date 16-feb-1940 -location campbeltown -ref isabella1850-death
        partner {
            name Peter McGregor
            gender male
        }
    }
    child {
        name Margaret Galbraith
        gender female
        baptism -date 29-aug-1853 -location campbeltown -ref margaret1853-baptism
        death -date 31-mar-1857 -location campletown -ref margaret1853-death
    }
    child galbraith-donald-1856-mcpherson

    child {
        name Alexander Galbreath
        gender male
        birth -date 13-may-1859 -location campbeltown -ref alexander1859-birth
    }
}

footnotes {
    donald1819-baptism {
        $opr-ref[b-1819-519-000-0010-0305 "Donald Galbreath"]
    }
    death {
        $sp-ref-link[d-1897-507-00-0150 0050 "Donald Galbraith"]{
            150. Donald Galbraith, Cartwright (journeyman), widower of Isabella McLean;
            1897 November Twentieth, Age 66 [$i{abt 1831}]; 
            Father: Donald Galbraith, Customer Weaver[$i{illegible}] (deceased);
            Mother: Annie Galbraith, M.S. Bell (deceased);
            Present: Donald Galbraith. son
        }
        The age of 66 is incorrect.  That would mean he got married at age 16.
    }
    mclean-marriage {
        $opr-ref[m-1847-507-000-0060-0423 "Donald Galbreath" "Isabella McLean"]{
            Donald Galbreath Wright[?] Dalintober and Isabella McLean
            Daughter of the late Donald McLean Farmer Ballywilone[?]
            both of this Parish were married thirteenth December 1847.
        }
    }
    mclean-death {
        $sp-ref-link[d-1887-507-00-0177 0059 "Isabella Galbraith"]
    }
    agnes1848-baptism {
        $opr-ref[b-1848-507-000-0070-0330 "Agnes Galbreath"]
    }
    isabella1850-baptism {
        $opr-ref[b-1850-507-000-0070-0373 "Isabella Galbraith"]
    }
    isabella1850-death {
        $sp-ref-link[d-1940-507-00-0022 0008 "Isabella Galbraith"]
    }
    margaret1853-baptism {
        $opr-ref[b-1853-507-000-0070-0419 "Margaret Galbraith"]
    }
    margaret1853-death {
        $sp-ref-link[d-1857-507-00-0045 0015 "Margaret Galbraith"]
    }
    alexander1859-birth {
        $sp-ref-link[b-1859-507-00-0082 0028 "Alexander Galbreath"]
    }
    census1881 {
       https://www.findmypast.com/transcript?id=GBC%2F1881%2F0029341604&expand=true&tab=this
    }
}

