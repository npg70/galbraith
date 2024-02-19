name Donald Galbreath
birth -date 'about 1785'
tags gigha
external {
    familysearch KD3X-J7F
}
todo {
    Death
}
todo {
    check if the Double Mary is two people or one.  Baptisms are exactly one year apart.
}
todo {
    in 1851 Census, two grand sons are listed.   The only two that make sense are perhaps the family of $child-link[galbraith-archibald-1800-wilkinson]{Archibald Galbraith} and Barbara Wilkenson.
}

partner {
    name Catharine Brown
    birth -date 'about 1785'
    death -date 05-mar-1853 -location gigha

    marriage -date 15-dec-1802 -location gigha -ref brown-marriage

    child {
        name Margaret Galbreath
        baptism -date 14-jul-1803 -location gigha -ref margaret1803-baptism
        death -date 27-sep-1827 -location gigha -ref margaret1803-death
    }
    child galbreath-john-1805-milloy

    child {
        name Mary Galbreath
        baptism -date 01-nov-1807 -location gigha -ref mary1807-baptism
    }
    child {
        name Mary Galbreath
        baptism -date 01-nov-1808 -location gigha -ref mary1808-baptism
    }
    child {
        name Neil Galbreath
        baptism -date 06-jun-1810 -location gigha -ref neil1810-baptism
        death -date 10-sep-1810 -location gigha -ref neil1810-death
    }
    child {
        name Merran Galbreath
        baptism -date 14-nov-1811 -location gigha -ref merran1811-baptism
    }
    child {
        name Donald Galbreath
        baptism -date 18-oct-1814 -location gigha -ref donald1814-baptism
    }
    child {
        name Ann Galbreath
        birth -date 'about 1815' -location gigha
        death -date 08-jun-1835 -location gigha  -ref ann1815-death
    }
    child {
        name Catherine Galbreath
        birth -date 13-apr-1919  -location gigha -ref catherine1819-baptism
    }
    child {
        name Lachlan Galbreath
        baptism -date 24-aug-1822 -location gigha -ref lachlan1822-baptism
        death -date 05-jan-1829 -location gigha -ref lachlan1822-death
    }
}
body {

The 1827 Census of Gigha, at Kill (Keill): $ref[census1827]
$csvtable{
First,Last,Age,Year
Donald,Galbraith,42,1785
Catharine,Brown,42,1785
Peggy,Galbraith,22,1802
Mary,Galbraith,20,1804
More,Galbraith,18,1806
Donald,-,13,1814
Ann,-,10,1817
Catharine,-,5,1822
Lachlan,-,5,1822
}

The 1851 Census at Ardmeanish, Gigha, Kintyre, Argyllshire, Scotland:$ref[census1851]
$csvtable{
First,Last,Role,Age,Year,Occupation
Donald,Galbreath,Head,70,1781,Shoemaker
Catherine,Galbreath,Wife,71,1780,-
John,Galbreath,Son,45,1806,Shoemaker
Duncan,Galbreath,Grand son,16,1835,Scholar
Archibald,Galbreath,Grand son,7,1844,Scholar
}
All listed as being born in Gigha.

}
footnotes {

    brown-marriage {

    }
    margaret1803-baptism {
        $opr-ref[b-1803-537-000-0010-0026 "Margarate Galbreath"]
    }
    margaret1803-death {
        $opr-ref-link[d-1827-537-000-0020-0018 "Margarat Galbreath"]
    }
    mary1807-baptism {
        $opr-ref[b-1807-537-000-0010-0035 "Mary Galbreath"]
    }
    mary1808-baptism {
        $opr-ref[b-1808-537-000-0010-0037 "Mary Galbreath"]
    }
    neil1810-baptism {
        $opr-ref[b-1810-537-000-0010-0044 "Neil Galbreath"]
    }
    neil1810-death {
        $opr-ref[d-1810-537-000-0010-0045 "Neil Galbreath"]
    }
    merran1811-baptism {
        $opr-ref[b-1811-537-000-0010-0047 "Merran Galbreath"]
    }
    donald1814-baptism {
        $opr-ref-link[b-1814-537-000-0010-0052 "Donald Galbreath"]
    }
    ann1815-death {
        $opr-ref-link[d-1835-537-000-0020-0032 "Ann Galbreath"]{
            July | 8 | Ann Galbreath at Ardmish, age 20 years.
        }
    }
    catherine1819-baptism {
        $opr-ref[b-1819-537-000-0010-0061 "Catharine Galbreath"]
    }
    lachlan1822-baptism {
        $opr-ref-link[b-1822-537-000-0020-0007 "Lachlan Galbreath"]
    }
    lachlan1822-death {
        $opr-ref-link[d-1829-537-000-0020-0020 "Lachlan Galbreath"]
    }
    census1827 {
        from https://www.gigha.org.uk/viewItem.php?id=8932&parentId=8926&sectionTitle=About+Gigha
    }
    census1851 {
        https://www.findmypast.com/transcript?id=GBC/1851/0019321713
    }
}
