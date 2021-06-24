package src

func GetTopPorts(number int) string{

    var top_ports string
    switch number{
        case 100:
            return "7,9,13,21,22,23,25,26,37,53,79,80,81,88,106,110,111,113,119,135,139,143,144,179,199,389,427,443,444,445,465,513,514,515,543,544,548,554,587,631,646,873,990,993,995,1025,1026,1027,1028,1029,1110,1433,1720,1723,1755,1900,2000,2001,2049,2121,2717,3000,3128,3306,3389,3986,4899,5000,5009,5051,5060,5101,5190,5357,5432,5631,5666,5800,5900,6000,6001,6646,7070,8000,8008,8009,8080,8081,8443,8888,9100,9999,10000,32768,49152,49153,49156,49157" 
            break
       case 500:
           return "1,3,6,9,13,17,19,20,21,22,23,24,25,30,33,37,42,49,53,79,80,81,82,85,88,90,100,106,110,113,119,135,139,143,146,161,163,179,199,211,222,254,264,280,306,311,340,366,389,407,427,443,444,464,497,500,512,513,514,541,543,548,554,563,587,593,625,631,636,646,648,705,711,787,800,808,873,880,888,900,901,912,987,990,992,995,999,1002,1022,1023,1024,1025,1026,1027,1028,1029,1030,1031,1032,1033,1034,1035,1036,1037,1038,1039,1040,1041,1042,1043,1044,1045,1046,1047,1048,1049,1050,1051,1052,1053,1054,1055,1056,1057,1058,1059,1060,1061,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1077,1078,1079,1080,1081,1082,1085,1088,1093,1096,1097,1098,1099,1104,1106,1107,1110,1148,1169,1218,1234,1248,1272,1310,1352,1433,1494,1500,1503,1521,1666,1687,1700,1717,1720,1723,1755,1761,1783,1801,1840,1863,1900,1935,1947,1998,2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2030,2049,2065,2100,2103,2105,2107,2119,2121,2135,2144,2160,2190,2222,2260,2301,2381,2383,2399,2401,2492,2500,2525,2601,2604,2607,2701,2717,2809,2811,2869,2875,2967,3000,3005,3017,3031,3052,3071,3128,3211,3260,3268,3283,3300,3306,3323,3325,3333,3351,3367,3389,3404,3476,3551,3580,3659,3689,3703,3766,3784,3801,3827,3986,3998,4000,4001,4002,4045,4126,4129,4242,4443,4449,4662,4899,5000,5001,5002,5003,5009,5030,5050,5060,5100,5101,5120,5190,5214,5222,5225,5269,5357,5414,5431,5500,5550,5555,5566,5631,5633,5666,5679,5718,5800,5810,5825,5877,5900,5901,5910,5925,5959,5960,5961,5987,5988,6000,6001,6004,6059,6101,6112,6123,6129,6156,6389,6543,6580,6646,6666,6788,6881,6901,6969,7000,7019,7070,7100,7106,7200,7625,7627,7741,7777,7911,7937,8000,8001,8007,8008,8009,8021,8031,8080,8081,8082,8083,8084,8085,8086,8087,8088,8181,8192,8193,8222,8291,8333,8400,8402,8443,8600,8649,8651,8701,8873,8888,8899,8994,9000,9001,9009,9050,9071,9090,9100,9101,9207,9415,9535,9593,9594,9876,9999,10000,10010,10243,12000,12345,13782,14238,15000,16992,20005,20828,23502,27000,32768,32769,32770,32771,32772,32773,32774,33354,35500,42510,45100,49152,49153,49154,49155,49156,49999,50000,50001,51103,52822,52869,55555,55600,64623,64680,65000,65389"
           break
      case 1000:
          return "1,3,6,9,13,17,19,20,21,22,23,24,25,30,32,37,42,49,53,70,79,80,81,82,83,84,88,89,99,106,109,110,113,119,125,135,139,143,146,161,163,179,199,211,222,254,255,259,264,280,301,306,311,340,366,389,406,416,425,427,443,444,458,464,481,497,500,512,513,514,524,541,543,544,548,554,563,587,593,616,625,631,636,646,648,666,667,683,687,691,700,705,711,714,720,722,726,749,765,777,783,787,800,808,843,873,880,888,898,900,901,902,911,981,987,990,992,995,999,1000,1001,1007,1009,1010,1021,1022,1023,1024,1025,1026,1027,1028,1029,1030,1031,1032,1033,1034,1035,1036,1037,1038,1039,1040,1041,1042,1043,1044,1045,1046,1047,1048,1049,1050,1051,1052,1053,1054,1055,1056,1057,1058,1059,1060,1061,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1086,1087,1088,1089,1090,1091,1092,1093,1094,1095,1096,1097,1098,1099,1102,1104,1105,1106,1107,1110,1111,1112,1113,1117,1119,1121,1122,1123,1126,1130,1131,1137,1141,1145,1147,1148,1151,1154,1163,1164,1165,1169,1174,1183,1185,1186,1192,1198,1201,1213,1216,1217,1233,1236,1244,1247,1259,1271,1277,1287,1296,1300,1309,1310,1322,1328,1334,1352,1417,1433,1443,1455,1461,1494,1500,1503,1521,1524,1533,1556,1580,1583,1594,1600,1641,1658,1666,1687,1700,1717,1718,1719,1720,1723,1755,1761,1782,1801,1805,1812,1839,1862,1863,1875,1900,1914,1935,1947,1971,1974,1984,1998,1999,2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2013,2020,2021,2030,2033,2034,2038,2040,2041,2042,2045,2046,2047,2048,2065,2068,2099,2103,2105,2106,2111,2119,2121,2126,2135,2144,2160,2170,2179,2190,2196,2200,2222,2251,2260,2288,2301,2323,2366,2381,2382,2393,2399,2401,2492,2500,2522,2525,2557,2601,2604,2607,2638,2701,2710,2717,2725,2800,2809,2811,2869,2875,2909,2920,2967,2998,3000,3003,3005,3006,3011,3013,3017,3030,3052,3071,3077,3128,3168,3211,3221,3260,3268,3283,3300,3306,3322,3323,3324,3333,3351,3367,3369,3370,3371,3389,3404,3476,3493,3517,3527,3546,3551,3580,3659,3689,3703,3737,3766,3784,3800,3809,3814,3826,3827,3851,3869,3871,3878,3880,3889,3905,3914,3918,3920,3945,3971,3986,3995,3998,4000,4001,4002,4003,4004,4005,4045,4111,4125,4129,4224,4242,4279,4321,4343,4443,4444,4445,4449,4550,4567,4662,4848,4899,4998,5000,5001,5002,5003,5009,5030,5033,5050,5054,5060,5080,5087,5100,5101,5120,5190,5200,5214,5221,5225,5269,5280,5298,5357,5405,5414,5431,5440,5500,5510,5544,5550,5555,5560,5566,5631,5633,5666,5678,5718,5730,5800,5801,5810,5815,5822,5825,5850,5859,5862,5877,5900,5901,5902,5903,5906,5910,5915,5922,5925,5950,5952,5959,5960,5961,5962,5987,5988,5998,5999,6000,6001,6002,6003,6004,6005,6006,6009,6025,6059,6100,6106,6112,6123,6129,6156,6346,6389,6502,6510,6543,6547,6565,6566,6580,6646,6666,6667,6668,6689,6692,6699,6779,6788,6792,6839,6881,6901,6969,7000,7001,7004,7007,7019,7025,7070,7100,7103,7106,7200,7402,7435,7443,7496,7512,7625,7627,7676,7741,7777,7800,7911,7920,7937,7999,8000,8001,8007,8008,8009,8010,8021,8031,8042,8045,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8093,8099,8180,8192,8193,8200,8222,8254,8290,8291,8300,8333,8383,8400,8402,8443,8500,8600,8649,8651,8654,8701,8800,8873,8888,8899,8994,9000,9001,9002,9009,9010,9040,9050,9071,9080,9090,9099,9100,9101,9102,9110,9200,9207,9220,9290,9415,9418,9485,9500,9502,9535,9575,9593,9594,9618,9666,9876,9877,9898,9900,9917,9929,9943,9968,9998,9999,10000,10001,10002,10003,10009,10012,10024,10082,10180,10215,10243,10566,10616,10621,10626,10628,10778,11110,11967,12000,12174,12265,12345,13456,13722,13782,14000,14238,14441,15000,15002,15003,15660,15742,16000,16012,16016,16018,16080,16113,16992,17877,17988,18040,18101,18988,19101,19283,19315,19350,19780,19801,19842,20000,20005,20031,20221,20828,21571,22939,23502,24444,24800,25734,26214,27000,27352,27355,27715,28201,30000,30718,30951,31038,31337,32768,32769,32770,32771,32772,32773,32774,32775,32776,32777,32778,32779,32780,32781,32782,32783,32784,33354,33899,34571,34572,35500,38292,40193,40911,41511,42510,44176,44442,44501,45100,48080,49152,49153,49154,49155,49156,49157,49158,49159,49160,49163,49165,49167,49175,49400,49999,50000,50001,50002,50006,50300,50389,50500,50636,50800,51103,51493,52673,52822,52848,52869,54045,54328,55055,55555,55600,56737,57294,57797,58080,60020,60443,61532,61900,62078,63331,64623,64680,65000,65129,65389"
          break
      case 5000:
          return "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70,71,72,73,74,75,76,77,78,79,80,81,82,83,84,85,86,87,88,89,90,91,92,93,94,95,96,97,98,99,100,101,102,103,104,105,106,107,108,109,110,111,112,113,114,115,116,117,118,119,120,121,122,123,124,125,126,127,128,129,130,131,132,133,134,135,136,137,138,139,140,141,142,143,144,145,146,147,148,149,150,151,152,153,154,155,156,157,158,159,160,161,162,163,164,165,166,167,168,169,170,171,172,173,174,175,176,177,178,179,180,181,182,183,184,185,186,187,188,189,190,191,192,193,194,195,196,197,198,199,200,201,202,203,204,205,206,207,208,209,210,211,212,213,214,215,216,217,218,219,220,221,222,223,224,225,228,229,230,233,234,235,236,237,242,243,244,245,246,247,248,249,250,251,252,253,254,255,256,257,258,259,260,261,262,263,264,265,266,267,268,269,270,273,276,280,281,282,283,286,287,288,293,294,300,303,305,308,309,310,311,312,313,314,315,316,317,318,319,320,321,322,323,324,325,329,333,336,340,343,344,345,346,347,348,349,350,351,352,353,354,355,356,357,358,359,360,361,362,363,364,365,366,367,368,369,370,371,372,373,374,375,376,377,378,379,380,381,382,383,384,385,386,387,388,389,390,391,392,393,394,395,396,397,398,399,400,401,402,403,404,405,406,407,408,409,410,411,412,413,414,415,416,417,418,419,420,421,422,423,424,425,426,427,428,429,430,431,432,433,434,435,436,437,438,439,440,441,442,443,444,445,446,447,448,449,450,451,452,453,454,455,456,457,458,459,460,461,462,463,464,465,466,467,468,469,470,471,472,473,474,475,476,477,478,479,480,481,482,483,484,485,486,487,488,489,490,491,492,493,494,495,496,497,498,499,500,501,502,503,504,505,506,507,508,509,510,511,512,513,514,515,516,517,518,519,520,521,522,523,524,525,526,527,528,529,530,531,532,533,534,535,536,537,538,539,540,541,542,543,544,545,546,547,548,549,550,551,552,553,554,555,556,557,558,559,560,561,562,563,564,565,566,567,568,569,570,571,572,573,574,575,576,577,578,579,580,581,582,583,584,585,586,587,588,589,590,591,592,593,594,595,596,597,598,599,600,601,602,603,604,605,606,607,608,609,610,611,612,613,614,615,616,617,618,619,620,621,622,623,624,625,626,627,628,629,630,631,632,633,634,635,636,637,638,639,640,641,642,643,644,645,646,647,648,649,650,651,652,653,654,655,656,657,658,659,660,661,662,663,664,665,666,667,668,669,670,671,672,673,674,675,676,677,678,679,680,681,682,683,684,685,686,687,688,689,690,691,692,693,694,695,696,697,698,699,700,701,702,703,704,705,706,707,708,709,710,711,712,713,714,715,716,717,718,719,720,721,722,723,724,725,726,727,728,729,730,731,732,733,734,735,736,737,738,739,740,741,742,743,744,745,746,747,748,749,750,751,752,753,754,755,756,757,758,759,760,761,762,763,764,765,766,767,768,769,770,771,772,773,774,775,776,777,778,779,780,781,782,783,784,785,786,787,788,789,790,791,792,793,794,795,796,797,798,799,800,801,802,803,804,805,806,807,808,809,810,811,812,813,814,815,816,817,818,819,820,821,822,823,824,825,826,827,828,829,830,831,832,833,834,835,836,837,838,839,840,841,842,843,844,845,846,847,848,849,850,851,852,853,854,855,856,857,858,859,860,861,862,863,864,865,866,867,868,869,870,871,872,873,874,875,876,877,878,879,880,881,882,883,884,885,886,887,888,889,890,891,892,893,894,895,896,897,898,899,900,901,902,903,904,905,906,907,908,909,910,911,912,913,914,915,918,919,920,921,922,923,924,925,926,927,928,929,930,931,932,933,934,935,936,937,938,941,942,943,944,945,946,947,948,949,950,951,952,953,954,955,956,957,958,959,960,961,962,963,964,965,966,967,968,969,970,971,972,973,974,975,976,977,978,979,980,981,982,983,984,985,986,987,988,989,990,991,992,993,994,995,996,997,998,999,1000,1001,1002,1003,1004,1005,1006,1007,1008,1009,1010,1011,1012,1013,1014,1015,1016,1017,1018,1019,1020,1021,1022,1023,1024,1025,1026,1027,1028,1029,1030,1031,1032,1033,1034,1035,1036,1037,1038,1039,1040,1041,1042,1043,1044,1045,1046,1047,1048,1049,1050,1051,1052,1053,1054,1055,1056,1057,1058,1059,1060,1061,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1086,1087,1088,1089,1090,1091,1092,1093,1094,1095,1096,1097,1098,1099,1100,1101,1102,1103,1104,1105,1106,1107,1108,1109,1110,1111,1112,1113,1114,1115,1116,1117,1118,1119,1120,1121,1122,1123,1124,1125,1126,1127,1128,1129,1130,1131,1132,1133,1134,1135,1136,1137,1138,1139,1140,1141,1142,1143,1144,1145,1146,1147,1148,1149,1150,1151,1152,1153,1154,1155,1156,1157,1158,1159,1160,1161,1162,1163,1164,1165,1166,1167,1168,1169,1170,1171,1172,1173,1174,1175,1176,1177,1178,1179,1180,1181,1182,1183,1184,1185,1186,1187,1188,1189,1190,1191,1192,1193,1194,1195,1196,1197,1198,1199,1200,1201,1202,1203,1204,1205,1206,1207,1208,1209,1210,1211,1212,1213,1214,1215,1216,1217,1218,1219,1220,1221,1222,1223,1224,1225,1226,1227,1228,1229,1230,1231,1232,1233,1234,1235,1236,1237,1238,1239,1240,1241,1242,1243,1244,1245,1246,1247,1248,1249,1250,1251,1252,1253,1254,1255,1256,1257,1258,1259,1260,1261,1262,1263,1264,1265,1266,1267,1268,1269,1270,1271,1272,1273,1274,1275,1276,1277,1278,1279,1280,1281,1282,1283,1284,1285,1286,1287,1288,1289,1290,1291,1292,1293,1294,1295,1296,1297,1298,1299,1300,1301,1302,1303,1304,1305,1306,1307,1308,1309,1310,1311,1312,1313,1314,1315,1316,1317,1318,1319,1320,1321,1322,1323,1324,1325,1326,1327,1328,1329,1330,1331,1332,1333,1334,1335,1336,1337,1338,1339,1340,1341,1342,1343,1344,1345,1346,1347,1348,1349,1350,1351,1352,1353,1354,1355,1356,1357,1358,1359,1360,1361,1362,1363,1364,1365,1366,1367,1368,1369,1370,1371,1372,1373,1374,1375,1376,1377,1378,1379,1380,1381,1382,1383,1384,1385,1386,1387,1388,1389,1390,1391,1392,1393,1394,1395,1396,1397,1398,1399,1400,1401,1402,1403,1404,1405,1406,1407,1408,1409,1410,1411,1412,1413,1414,1415,1416,1417,1418,1419,1420,1421,1422,1423,1424,1425,1426,1427,1428,1429,1430,1431,1432,1433,1434,1435,1436,1437,1438,1439,1440,1441,1442,1443,1444,1445,1446,1447,1448,1449,1450,1451,1452,1453,1454,1455,1456,1457,1458,1459,1460,1461,1462,1463,1464,1465,1466,1467,1468,1469,1470,1471,1472,1473,1474,1475,1476,1477,1478,1479,1480,1481,1482,1483,1484,1485,1486,1487,1488,1489,1490,1491,1492,1493,1494,1495,1496,1497,1498,1499,1500,1501,1502,1503,1504,1505,1506,1507,1508,1509,1510,1511,1512,1513,1514,1515,1516,1517,1518,1519,1520,1521,1522,1523,1524,1525,1526,1527,1528,1529,1530,1531,1532,1533,1534,1535,1536,1537,1538,1539,1540,1541,1542,1543,1544,1545,1546,1547,1548,1549,1550,1551,1552,1553,1554,1555,1556,1557,1558,1559,1560,1561,1562,1563,1564,1565,1566,1567,1568,1569,1570,1571,1572,1573,1574,1575,1576,1577,1578,1579,1580,1581,1582,1583,1584,1585,1586,1587,1588,1589,1590,1591,1592,1593,1594,1595,1596,1597,1598,1599,1600,1601,1602,1603,1604,1605,1606,1607,1608,1609,1610,1611,1612,1613,1614,1615,1616,1617,1618,1619,1620,1621,1622,1623,1624,1625,1626,1627,1628,1629,1630,1631,1632,1633,1634,1635,1636,1637,1638,1639,1640,1641,1642,1643,1644,1645,1646,1647,1648,1649,1650,1651,1652,1653,1654,1655,1656,1657,1658,1659,1660,1661,1662,1663,1664,1665,1666,1667,1668,1669,1670,1671,1672,1673,1674,1675,1676,1677,1678,1679,1680,1681,1682,1683,1684,1685,1686,1687,1688,1689,1690,1691,1692,1693,1694,1695,1696,1697,1698,1699,1700,1701,1702,1703,1704,1705,1706,1707,1708,1709,1710,1711,1712,1713,1714,1715,1716,1717,1718,1719,1720,1721,1722,1723,1724,1725,1726,1727,1728,1729,1730,1731,1732,1733,1734,1735,1736,1737,1738,1739,1740,1741,1742,1743,1744,1745,1746,1747,1748,1749,1750,1751,1752,1753,1754,1755,1756,1757,1758,1759,1760,1761,1762,1763,1764,1765,1766,1767,1768,1769,1770,1771,1772,1773,1774,1775,1776,1777,1778,1779,1780,1781,1782,1783,1784,1785,1786,1787,1788,1789,1790,1791,1792,1793,1794,1795,1796,1797,1798,1799,1800,1801,1802,1803,1804,1805,1806,1807,1808,1809,1810,1811,1812,1813,1814,1815,1816,1817,1818,1819,1820,1821,1822,1823,1824,1825,1826,1827,1828,1829,1830,1831,1832,1833,1834,1835,1836,1837,1838,1839,1840,1841,1842,1843,1844,1845,1846,1847,1848,1849,1850,1851,1852,1853,1854,1855,1856,1857,1858,1859,1860,1861,1862,1863,1864,1865,1866,1867,1868,1869,1870,1871,1872,1873,1874,1875,1876,1877,1878,1879,1880,1881,1882,1883,1884,1885,1886,1887,1888,1889,1890,1891,1892,1893,1896,1897,1898,1899,1900,1901,1902,1903,1904,1905,1906,1907,1908,1909,1910,1911,1912,1913,1914,1915,1916,1917,1918,1919,1920,1921,1922,1923,1924,1925,1926,1927,1928,1929,1930,1931,1932,1933,1934,1935,1936,1937,1938,1939,1940,1941,1942,1943,1944,1945,1946,1947,1948,1949,1950,1951,1952,1953,1954,1955,1956,1957,1958,1959,1960,1961,1962,1963,1964,1965,1966,1967,1968,1969,1970,1971,1972,1973,1974,1975,1976,1977,1978,1979,1980,1981,1982,1983,1984,1985,1986,1987,1988,1989,1990,1991,1992,1993,1994,1995,1996,1997,1998,1999,2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022,2023,2024,2025,2026,2027,2028,2029,2030,2031,2032,2033,2034,2035,2036,2037,2038,2039,2040,2041,2042,2043,2044,2045,2046,2047,2048,2049,2050,2051,2052,2053,2054,2055,2056,2057,2058,2059,2060,2061,2062,2063,2064,2065,2066,2067,2068,2069,2070,2071,2072,2073,2074,2075,2076,2077,2078,2079,2080,2081,2082,2083,2084,2085,2086,2087,2088,2089,2090,2091,2092,2093,2094,2095,2096,2097,2098,2099,2100,2101,2102,2103,2104,2105,2106,2107,2108,2109,2110,2111,2112,2113,2114,2115,2116,2117,2118,2119,2120,2121,2122,2123,2124,2125,2126,2127,2128,2129,2130,2131,2132,2133,2134,2135,2136,2137,2138,2139,2140,2141,2142,2143,2144,2145,2146,2147,2148,2149,2150,2151,2152,2153,2154,2155,2156,2157,2158,2159,2160,2161,2162,2163,2164,2165,2166,2167,2168,2169,2170,2171,2172,2173,2174,2175,2176,2177,2178,2179,2180,2181,2182,2183,2184,2185,2186,2187,2188,2189,2190,2191,2192,2196,2197,2198,2199,2200,2201,2202,2203,2204,2205,2206,2207,2208,2209,2210,2211,2212,2213,2214,2215,2216,2217,2218,2219,2220,2221,2222,2223,2224,2232,2241,2250,2253,2260,2261,2265,2269,2270,2280,2288,2291,2300,2301,2304,2307,2312,2323,2325,2330,2335,2340,2366,2371,2375,2381,2382,2391,2393,2399,2401,2418,2425,2430,2431,2432,2435,2438,2449,2456,2463,2472,2492,2500,2505,2522,2525,2531,2550,2557,2564,2567,2580,2583,2598,2600,2601,2604,2605,2606,2607,2622,2627,2631,2638,2644,2691,2700,2701,2706,2710,2711,2717,2723,2725,2728,2734,2766,2800,2804,2806,2809,2811,2847,2850,2869,2875,2882,2888,2898,2901,2902,2908,2909,2920,2930,2957,2967,2973,2984,2987,2991,2997,3000,3001,3002,3005,3006,3011,3013,3017,3023,3025,3030,3045,3049,3052,3057,3062,3063,3071,3077,3080,3086,3089,3102,3118,3121,3128,3141,3146,3162,3167,3190,3200,3210,3220,3240,3260,3263,3268,3280,3283,3291,3299,3300,3304,3306,3310,3319,3322,3323,3324,3333,3351,3362,3365,3367,3368,3369,3370,3371,3374,3376,3388,3389,3396,3397,3398,3399,3404,3410,3414,3419,3421,3425,3430,3439,3443,3456,3476,3479,3483,3485,3493,3497,3503,3505,3511,3513,3514,3517,3519,3526,3530,3531,3546,3551,3577,3580,3586,3599,3602,3621,3632,3636,3652,3656,3658,3663,3669,3672,3680,3683,3689,3697,3700,3703,3712,3728,3731,3737,3742,3749,3765,3784,3787,3790,3792,3795,3798,3799,3800,3803,3806,3808,3809,3810,3811,3812,3813,3817,3820,3823,3824,3825,3826,3827,3830,3837,3839,3842,3846,3847,3848,3849,3850,3851,3852,3856,3859,3863,3868,3869,3870,3871,3876,3878,3879,3882,3888,3889,3897,3899,3900,3901,3904,3905,3906,3907,3908,3911,3913,3914,3915,3918,3919,3922,3928,3929,3930,3935,3936,3940,3943,3944,3945,3948,3952,3956,3961,3962,3963,3967,3968,3971,3975,3979,3980,3981,3982,3983,3984,3985,3989,3990,3991,3992,3993,3994,3995,3996,3997,3998,3999,4000,4001,4002,4003,4004,4005,4006,4007,4008,4009,4016,4020,4022,4024,4029,4035,4039,4045,4056,4058,4065,4080,4087,4090,4096,4100,4111,4112,4118,4119,4120,4125,4129,4132,4135,4141,4143,4147,4158,4161,4164,4174,4190,4192,4199,4206,4220,4224,4234,4242,4252,4262,4279,4294,4297,4300,4302,4321,4325,4328,4333,4342,4355,4356,4357,4369,4374,4375,4384,4388,4401,4407,4414,4418,4430,4433,4442,4443,4444,4445,4446,4449,4454,4464,4471,4476,4480,4500,4516,4530,4534,4545,4550,4555,4557,4558,4567,4570,4599,4600,4601,4606,4609,4644,4649,4658,4660,4662,4665,4672,4687,4689,4700,4712,4745,4760,4767,4770,4778,4793,4800,4819,4848,4859,4875,4876,4881,4899,4903,4912,4931,4949,4987,4998,4999,5000,5001,5002,5003,5004,5009,5010,5011,5012,5013,5014,5015,5016,5020,5023,5030,5033,5040,5050,5051,5052,5053,5054,5060,5063,5066,5070,5074,5080,5087,5090,5095,5098,5100,5101,5111,5114,5120,5121,5125,5133,5137,5145,5147,5151,5190,5193,5200,5201,5212,5214,5219,5221,5222,5225,5232,5233,5234,5242,5250,5252,5259,5261,5269,5279,5291,5298,5300,5301,5302,5308,5339,5347,5353,5357,5370,5377,5400,5405,5414,5423,5431,5432,5440,5441,5444,5457,5473,5475,5490,5500,5501,5510,5520,5530,5544,5550,5552,5553,5554,5557,5560,5566,5580,5611,5620,5621,5631,5632,5665,5666,5672,5678,5679,5711,5713,5717,5721,5722,5730,5732,5734,5737,5800,5801,5802,5803,5806,5807,5810,5811,5814,5817,5820,5821,5822,5823,5824,5825,5826,5831,5834,5836,5838,5839,5845,5848,5849,5852,5853,5858,5859,5862,5868,5871,5874,5877,5881,5887,5899,5900,5901,5902,5903,5904,5905,5906,5907,5908,5909,5910,5911,5914,5917,5920,5921,5922,5923,5924,5925,5926,5931,5934,5936,5938,5939,5945,5948,5949,5952,5953,5958,5959,5960,5961,5962,5966,5968,5971,5974,5977,5981,5985,5986,5987,5988,5997,5998,5999,6000,6001,6002,6003,6004,6005,6006,6007,6008,6009,6015,6017,6021,6025,6030,6050,6051,6055,6059,6062,6065,6067,6077,6085,6090,6100,6103,6105,6110,6111,6112,6115,6120,6123,6126,6129,6141,6142,6145,6146,6156,6161,6203,6222,6247,6250,6259,6273,6309,6323,6346,6349,6379,6389,6400,6412,6481,6500,6502,6503,6510,6520,6535,6543,6547,6550,6565,6566,6579,6588,6600,6606,6619,6628,6644,6646,6650,6662,6665,6666,6667,6668,6669,6689,6692,6699,6700,6709,6710,6725,6732,6734,6779,6788,6792,6839,6877,6881,6888,6896,6901,6920,6922,6942,6956,6969,6972,7000,7001,7002,7003,7004,7005,7006,7007,7008,7009,7019,7024,7033,7043,7050,7067,7070,7071,7080,7092,7099,7100,7101,7102,7103,7106,7119,7121,7123,7184,7200,7218,7231,7241,7272,7278,7281,7300,7320,7325,7345,7400,7402,7435,7438,7443,7451,7456,7464,7496,7500,7512,7553,7555,7597,7600,7625,7627,7634,7637,7654,7676,7685,7688,7725,7741,7744,7749,7770,7771,7777,7780,7788,7800,7813,7830,7852,7853,7878,7895,7900,7911,7913,7920,7929,7937,7975,7998,7999,8000,8001,8002,8005,8006,8007,8008,8009,8010,8014,8015,8018,8021,8022,8025,8029,8031,8037,8042,8045,8050,8052,8060,8064,8076,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8092,8095,8097,8098,8099,8110,8116,8118,8123,8133,8144,8180,8189,8192,8193,8200,8201,8222,8232,8245,8248,8254,8268,8273,8282,8290,8291,8292,8293,8294,8300,8308,8333,8339,8383,8385,8400,8401,8402,8409,8443,8445,8451,8452,8453,8454,8470,8471,8474,8477,8479,8481,8484,8500,8515,8530,8539,8562,8600,8621,8640,8644,8648,8651,8654,8658,8673,8675,8680,8686,8701,8736,8752,8756,8765,8770,8772,8790,8798,8800,8843,8865,8873,8877,8878,8879,8882,8887,8888,8892,8898,8899,8925,8954,8980,8987,8994,8996,8999,9000,9001,9002,9003,9004,9009,9010,9013,9020,9021,9037,9040,9044,9050,9061,9065,9071,9080,9084,9090,9098,9099,9100,9101,9102,9103,9104,9105,9106,9110,9125,9128,9130,9133,9152,9160,9170,9183,9191,9197,9200,9202,9207,9210,9220,9287,9290,9300,9333,9343,9351,9364,9400,9409,9415,9418,9443,9454,9464,9478,9485,9493,9500,9501,9502,9513,9527,9535,9575,9583,9592,9593,9594,9600,9613,9616,9618,9619,9620,9628,9643,9648,9654,9661,9665,9666,9673,9679,9683,9694,9700,9745,9777,9812,9814,9823,9825,9830,9844,9875,9876,9877,9898,9900,9908,9909,9910,9911,9914,9917,9919,9929,9941,9943,9950,9968,9971,9975,9979,9988,9990,9991,9995,9998,9999,10000,10001,10002,10003,10004,10005,10006,10007,10008,10009,10010,10011,10018,10019,10022,10023,10024,10034,10042,10045,10058,10064,10082,10093,10101,10115,10160,10180,10215,10238,10243,10245,10255,10280,10338,10347,10357,10387,10414,10443,10494,10500,10509,10529,10535,10550,10551,10552,10553,10554,10555,10565,10566,10601,10616,10621,10626,10628,10699,10754,10778,10842,10852,10873,10878,10900,11000,11003,11007,11019,11026,11031,11032,11089,11100,11110,11180,11200,11224,11250,11288,11296,11371,11401,11552,11697,11735,11813,11862,11940,11967,12000,12001,12005,12009,12019,12021,12031,12034,12059,12077,12080,12090,12096,12121,12132,12137,12146,12156,12171,12174,12192,12215,12225,12240,12243,12251,12262,12265,12271,12275,12296,12340,12345,12380,12414,12452,12699,12702,12766,12865,12891,12955,12962,13017,13093,13130,13132,13140,13142,13149,13167,13188,13192,13193,13229,13250,13261,13264,13306,13318,13340,13359,13456,13502,13580,13695,13701,13713,13714,13718,13720,13721,13722,13723,13730,13766,13782,13783,13846,13899,14000,14141,14147,14218,14237,14254,14418,14441,14442,14443,14534,14545,14693,14733,14827,14891,14916,15000,15001,15002,15003,15004,15050,15145,15151,15190,15275,15317,15344,15402,15448,15550,15631,15645,15660,15670,15677,15722,15730,15742,15758,15915,16000,16012,16016,16018,16048,16080,16113,16161,16270,16273,16283,16286,16297,16349,16372,16444,16464,16705,16723,16724,16797,16800,16845,16851,16900,16992,17007,17016,17070,17089,17129,17251,17255,17300,17409,17413,17500,17595,17700,17701,17715,17801,17860,17867,17877,17969,17985,17988,17997,18000,18012,18015,18018,18040,18080,18101,18148,18181,18182,18183,18187,18231,18264,18333,18336,18380,18439,18505,18517,18569,18669,18874,18887,18910,18962,18988,19010,19101,19130,19150,19200,19283,19315,19333,19350,19353,19403,19464,19501,19612,19634,19715,19780,19801,19842,19852,19900,19995,20000,20001,20005,20011,20017,20021,20031,20039,20052,20076,20080,20085,20089,20102,20106,20111,20118,20125,20127,20147,20179,20221,20222,20223,20224,20225,20226,20227,20280,20473,20734,20828,20883,20934,20940,20990,21011,21078,21201,21473,21571,21631,21634,21728,21792,21891,21915,22022,22063,22100,22125,22128,22177,22200,22222,22273,22290,22341,22350,22555,22563,22711,22719,22727,22769,22882,22939,22959,22969,23017,23040,23052,23219,23228,23270,23296,23342,23382,23430,23451,23502,23723,23796,23887,23953,24218,24392,24416,24444,24552,24554,24616,24800,24999,25000,25174,25260,25262,25288,25327,25445,25473,25486,25565,25703,25717,25734,25847,26000,26007,26208,26214,26340,26417,26470,26669,26972,27000,27001,27002,27005,27007,27009,27015,27016,27017,27018,27055,27074,27087,27204,27316,27350,27351,27352,27355,27356,27372,27374,27521,27537,27665,27715,27770,28017,28114,28142,28201,28211,28374,28567,28717,28850,28924,28967,29045,29152,29243,29507,29672,29810,29831,30000,30005,30087,30195,30299,30519,30599,30644,30659,30704,30718,30896,30951,31033,31038,31058,31072,31337,31339,31386,31416,31438,31522,31657,31727,32006,32022,32031,32088,32102,32200,32219,32260,32764,32767,32768,32769,32770,32771,32772,32773,32774,32775,32776,32777,32778,32779,32780,32781,32782,32783,32784,32785,32786,32787,32788,32789,32790,32791,32797,32798,32803,32807,32814,32815,32820,32822,32835,32837,32842,32858,32868,32871,32888,32897,32904,32908,32910,32932,32944,32960,32976,33000,33011,33017,33070,33087,33124,33175,33192,33200,33203,33277,33327,33335,33337,33354,33367,33395,33444,33453,33522,33550,33554,33604,33841,33879,33882,33889,33895,33899,34021,34036,34096,34189,34317,34341,34381,34401,34507,34510,34571,34572,34683,34728,34765,34783,34833,34875,35033,35050,35116,35131,35217,35272,35349,35392,35401,35500,35506,35513,35553,35593,35731,35879,35900,35906,35929,35986,36046,36104,36256,36275,36368,36436,36508,36530,36552,36659,36677,36694,36710,36748,36823,36914,36950,36962,36983,37121,37151,37174,37185,37218,37393,37522,37607,37614,37647,37674,37777,37789,37839,37855,38029,38037,38185,38188,38194,38205,38224,38270,38292,38313,38331,38358,38446,38481,38546,38561,38570,38761,38764,38780,38805,38936,39067,39117,39136,39265,39293,39376,39380,39433,39482,39489,39630,39659,39732,39763,39774,39795,39869,39883,39895,39917,40000,40001,40002,40005,40011,40193,40306,40393,40400,40457,40489,40513,40614,40628,40712,40732,40754,40811,40834,40911,40951,41064,41123,41142,41250,41281,41318,41342,41345,41348,41398,41442,41511,41523,41551,41632,41773,41794,41808,42001,42035,42127,42158,42251,42276,42322,42449,42452,42510,42559,42575,42590,42632,42675,42679,42685,42735,42906,42990,43000,43002,43018,43027,43103,43139,43143,43188,43212,43231,43242,43425,43654,43690,43734,43823,43868,44004,44101,44119,44176,44200,44334,44380,44410,44431,44442,44479,44501,44505,44541,44616,44628,44704,44709,44711,44965,44981,45038,45050,45100,45136,45164,45220,45226,45413,45438,45463,45602,45624,45697,45777,45864,45960,46034,46069,46115,46171,46182,46200,46310,46372,46418,46436,46593,46813,46992,46996,47012,47029,47119,47197,47267,47348,47372,47448,47544,47557,47567,47581,47595,47624,47634,47700,47777,47806,47850,47858,47860,47966,47969,48009,48067,48080,48083,48127,48153,48167,48356,48434,48619,48631,48648,48682,48783,48813,48925,48966,48973,49002,49048,49132,49152,49153,49154,49155,49156,49157,49158,49159,49160,49163,49164,49165,49166,49167,49168,49169,49170,49171,49172,49175,49179,49186,49189,49190,49195,49196,49201,49202,49203,49211,49213,49216,49228,49232,49235,49241,49275,49302,49352,49372,49398,49400,49452,49498,49500,49519,49520,49521,49597,49603,49678,49751,49762,49765,49803,49927,49999,50000,50001,50002,50006,50016,50019,50040,50050,50101,50189,50198,50202,50205,50224,50246,50258,50277,50300,50356,50389,50500,50513,50529,50545,50576,50585,50636,50692,50733,50787,50800,50809,50815,50831,50833,50834,50835,50849,50854,50887,50903,50945,50997,51011,51020,51037,51067,51103,51118,51139,51191,51233,51234,51240,51300,51343,51351,51366,51413,51423,51460,51484,51488,51493,51515,51582,51658,51771,51800,51809,51906,51909,51961,51965,52000,52001,52002,52025,52046,52071,52173,52225,52230,52237,52262,52391,52477,52506,52573,52660,52665,52673,52675,52710,52735,52822,52847,52848,52849,52850,52853,52869,52893,52948,53085,53178,53189,53211,53240,53313,53319,53361,53370,53460,53469,53491,53535,53633,53639,53656,53690,53742,53782,53827,53852,53910,53958,54045,54075,54101,54127,54235,54263,54276,54320,54323,54328,54514,54551,54605,54658,54688,54722,54741,54873,54907,54987,54991,55000,55020,55055,55183,55187,55227,55312,55350,55382,55400,55426,55479,55527,55555,55568,55576,55579,55600,55635,55652,55684,55721,55758,55773,55781,55901,55907,55910,55948,56016,56055,56259,56293,56507,56535,56591,56668,56681,56723,56725,56737,56810,56822,56827,56973,56975,57020,57103,57123,57294,57325,57335,57347,57350,57352,57387,57398,57479,57576,57665,57678,57681,57702,57730,57733,57797,57891,57896,57923,57928,57988,57999,58001,58072,58080,58107,58109,58164,58252,58305,58310,58374,58430,58446,58456,58468,58498,58562,58570,58610,58622,58630,58632,58634,58699,58721,58838,58908,58970,58991,59087,59107,59110,59122,59149,59160,59191,59200,59201,59239,59340,59499,59504,59509,59525,59565,59684,59778,59810,59829,59841,59987,60000,60002,60020,60055,60086,60111,60123,60146,60177,60227,60243,60279,60377,60401,60403,60443,60485,60492,60504,60544,60579,60612,60621,60628,60642,60713,60728,60743,60753,60782,60789,60794,60989,61159,61169,61402,61473,61516,61532,61613,61616,61669,61722,61734,61827,61851,61900,61942,62006,62042,62078,62080,62188,62312,62519,62570,62674,62866,63105,63156,63331,63423,63675,63803,64080,64127,64320,64438,64507,64551,64623,64680,64726,64890,65000,65048,65129,65301,65310,65389,65488,65514"
          break
        default:
            return "7,9,13,21,22,23,25,26,37,53,79,80,81,88,106,110,111,113,119,135,139,143,144,179,199,389,427,443,444,445,465,513,514,515,543,544,548,554,587,631,646,873,990,993,995,1025,1026,1027,1028,1029,1110,1433,1720,1723,1755,1900,2000,2001,2049,2121,2717,3000,3128,3306,3389,3986,4899,5000,5009,5051,5060,5101,5190,5357,5432,5631,5666,5800,5900,6000,6001,6646,7070,8000,8008,8009,8080,8081,8443,8888,9100,9999,10000,32768,49152,49153,49156,49157" 
        }
    return top_ports
}
