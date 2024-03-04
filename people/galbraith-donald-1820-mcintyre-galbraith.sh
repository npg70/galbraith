name Donald Galbraith
tags argyll:islay argyll:campbeltown
external {
    familysearch M1TN-RQF
    findagrave 149818386
}

birth -date 15-may-1820 -location islay -note "not registered"
death -date 22-mar-1863 -location clachan,Kilcalmonell -ref donald1820-death

note {
    According to Donald's death records his parents were Lachlan Galbraith and Mary McArthur. But, TBD the father was listed as John Galbraith.
}
todo {
    shows up in 1861 census visiting $child-link[galbraith-angus-1827-smith]{Angus Galbraith}: https://www.findmypast.com/transcript?id=GBC/1861/0022346873&expand=true
}

body {
    He was Minister of the Campbeltown congregation, at later converted to be an independant minister.
}
partner {
    name Mary McIntyre
    birth -date 'about 1826' -location islay -note 'the daughter of Donald McIntyre'
    death -date 27-apr-1851 -location campbeltown -ref mcintyre-death
    marriage -date 27-mar-1850 -location campbeltown  -ref mcintyre-marriage1 -ref mcintyre-marriage2
}
partner {
    name Margaret McKillop Galbraith
    birth -date 2-oct-1821 -location cathcart,renfrewshire
    death -date 25-jan-1903 -location glasgow -ref galbraith1821-death
    marriage -date 7-jun-1855 -location glasgow -ref galbraith1821-marriage
    body {
        She was the daughter of $child-link[galbraith-james-1792-fraser]{James Galbraith} and Helen Fraser.
    }

    child galbraith-john-alexander-1858-findlay

    child {
        name James Lachlan Galbraith
        birth -date 25-mar-1860 -location campbeltown -ref james1860-birth
        death -date 13-jun-1923 -location edinburgh -ref james1860-death -note unmarried
    }
    child {
        name Charles Donald Galbraith
        birth -date 6-feb-1862 -location campbeltown -ref charles1862-birth
        death -date 8-jun-1888 -note "by suicide. Unmarried" -ref charles1862-death
        body {
            Buried with father.
        }
    }
}
footnotes {
    donald1820-death {
        $sp-ref-link[d-1863-516-01-0003 0001 "Donald Galbraith"]
    }
    galbraith1821-marriage {
        $sp-ref-link[m-1855-644-06-0178 0089 "Donald Galbraith" "Margaret M Galbraith"]
    }
    galbraith1821-death {
        $sp-ref[d-1903-646-03-0081 "Margaret McKillop Galbraith"]
    }
    james1860-birth {
        $sp-ref-link[b-1860-507-00-0085 0029 "James Lachlan Galbraith"]
    }
    james1860-death {
        $sp-ref[d-1923-644-12-0553 "James Lachlan Gabraith"]
    }
    charles1862-birth {
        $sp-ref-link[b-1862-507-00-0051 0017 "Charles Donald Galbraith"]
    }
    mcintyre-marriage1 {
        Indexed twice:
        $opr-ref[m-1850-540-000-0010-0357 "Donald Galbraith" "Mary McIntyre"];
        $opr-ref-link[m-1850-507-000-0060-0439 "Donald Galbraith" "Mary McIntyre"]
    }
    mcintyre-marriage2 {
        Greenock Advertiser, 9 Apr 1850:
        $blockquote{
            At Gartacharra, Islay on the 27th ult., by the Rev McLaurin, the Rev. Donald Galbraith, Campbelton, to Mary daughter of Donald McIntyre., Gartacharra.
        }
    }
    mcintyre-death {
        Greenock Advertiser, 6 May 1851:
        $blockquote{
        DEATHS - At Campbelton, on 27th ult, Mary McIntyre, wife of the Rev. D. Galbraith, Independent minister there.
        }
    }
}
