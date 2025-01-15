name Agnes Harvey
external {
    familysearch M1YZ-RWZ
    findagrave 158184603
    ancestry tree/12079504/person/282642876039
    wikitree Harvie-507
}
todo {
    FindAGrave inscription for Kilkerran needs help.
}
note {
    No biographical details in marriage and baptism records.
}
note {
    Son James is missing baptism record.
}

body {

Census of 1792, Kyllipoll

$csvtable{
Name,Age,Year
Robert Culbertson, 68, 1724
Ann Harvie, 67, 1725
John Culbertson, 30, 1762
David Culbertson, 14, 1778
}
}
baptism -date 16-may-1725 -location campbeltown
partner {
   name Robert Culbertson
   birth -date 2-dec-1722 -location campbeltown
   death -date 20-aug-1799 -location campbeltown
   marriage -date 07-feb-1748 -location campbeltown -ref culbertson-marriage

   child {
       name Janet Culbertson
       baptism -date 04-mar-1750 -location campbeltown -ref janet1750-baptism
   }
   child {
       name William Culbertson
       baptism -date 15-jul-1753 -location campbeltown -ref william1753-baptism
       partner {
           name Margaret Smith
           marriage -date 22-mar-1778 -location hamilton,lanark,scotland
       }
    }
    child {
        name Robert Culbertson
        baptism -date 23-jun-1755 -location campbeltown -ref robert1755-baptism
        death -note 'y'
    }
    child {
        name Robert Culbertson
        baptism -date 10-jan-1758 -location campbeltown -ref robert1758-baptism
    }
    child culbertson-james-1759-white

    child {
        name John Culbertson
        baptism -date 06-jun-1760 -location campbeltown -ref john1760-baptism
    }
}

footnotes {

    culbertson-marriage {
        Note: two other Galbraith marriages on the same page.
        $opr-ref-link[m-1748-507-000-0011-0450 "Agnes Harvie" "Robert Culbertson"]{
            Robert Culbertson & Agnes Harvie | Feb 7
        }
        $opr-ref[m-1748-507-000-0010-0187 "Robert Culbertson" "Agnes Harvie"]
    }
    janet1750-baptism {
        $opr-ref[b-1750-507-000-0010-0198 "Janet Culbertson"]
        $opr-ref-link[b-1750-507-000-0011-0271 "Janet Culbertson"]{
            Janet | Robert Culbertson and Agnes Harvie had a 
            Daughter baptized 4th March named Janet
        }
    }
    william1753-baptism {
        $opr-ref[b-1753-507-000-0011-0287 "William Culbertson"]
        $opr-ref-link[b-1753-507-000-0010-0213 "William Culbertson"]
    }
    robert1755-baptism {
        $opr-ref[b-1755-507-000-0010-0223 "Robert Culbertson"]
        $opr-ref[b-1755-507-000-0011-0296 "Robert Culbertson"]
    }
    robert1758-baptism {
        $opr-ref-link[b-1758-507-000-0010-0235 "Robert Culbertson"] 
        $opr-ref-link[b-1758-507-000-0011-0308 "Robert Culbertson"]{
            Robert | Robert Culbertson and Agnes Harvie had a son
            baptized 10th January named Robert
        }
    }      
    john1760-baptism {
        $opr-ref-link[b-1760-507-000-0011-0320 "John Culbertson"]{
            John | Robert Culbertson and Agnes Harvie had a son
            baptized 6th June named John.
        }
        $opr-ref-link[b-1760-507-000-0010-0249 "John Culbertson"]
    }
}
