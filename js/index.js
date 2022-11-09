window.onload = function () {
    let loadBox = document.getElementById("loading-box");
    loadBox.getElementsByClassName("loading-word")[0].textContent = "获取数据中，请稍后...";
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/GetData", true);
    xmlHttp.send(null);
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
            loadBox.getElementsByClassName("loading-word")[0].textContent = "解析数据中";
            FirstDraw(JSON.parse(xmlHttp.responseText));
        } else {
            document.getElementById("id");
        }
    }
}

function FirstDraw(data) {
    console.log(data);
    let root = document.getElementById("root");
    
    new Map(Object.entries(data["config"])).forEach(function (v, k) {
        let box = document.createElement('div');
        box.className="ShowData-Box";
        new Map(Object.entries(v)).forEach(function(arr,index){
            let level=document.createElement('div');
            level.className="ShowData-Level";
            arr.forEach(item=>{
                let btn=document.createElement('div');
                btn.className="ShowData-Btn";
                btn.textContent=item.Text;
                btn.setAttribute("ShowText",item.Text);
                btn.onclick=function(){
                    if(!btn.onCli){
                        btn.className="ShowData-Btn-cli";
                        btn.onCli=true;
                    }else{
                        btn.className="ShowData-Btn";
                        btn.onCli=false;
                    }
                };
                level.appendChild(btn);
            });
            box.appendChild(level);
        });
        root.appendChild(box);
    });
    let loadBox = document.getElementById("loading-box");
    setTimeout(function () {
        loadBox.className = "loaded";
        loadBox.style.display = 'none';
    }, 0);
}

// function MyVenn(){
//     let ss = ["奇数", "偶数", "质数", "2的幂", "3的倍数"];
//     let data = new Map();
//     data["奇数"] = [1, 3, 5, 7, 9];
//     data["偶数"] = [2, 4, 6, 8, 10];
//     data["质数"] = [2, 3, 5, 7];
//     data["2的幂"] = [2, 4, 8];
//     data["3的倍数"] = [3, 6, 9];
//     let lebels = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r"];
//     let sets = [];
//     let lebelmap = new Map();
//     let index = 0;
//     function ff(k) {
//         if (k.length < ss.length) {
//             for (let i = 0; i < 2; i++) {
//                 k.push(i);
//                 ff(k);
//                 k.pop();
//             }
//             return;
//         }
//         let name = [];
//         let lebel = "";
//         for (let i = 0; i < ss.length; i++) {
//             if (k[i] > 0) {
//                 name.push(ss[i]);
//                 lebel += ss[i];
//             }
//         }
//         if (name.length == 0) {
//             return;
//         }
//         let n = 0;
//         let nums = [];
//         for (let i = 1; i <= 10; i++) {
//             let flag = true;
//             for (let j = 0; j < ss.length; j++) {
//                 if (k[j] == 0)
//                     continue;
//                 if (data[ss[j]].indexOf(i) == -1) {
//                     flag = false;
//                     break;
//                 }
//             }
//             if (flag) {
//                 nums.push(i);
//                 n++;
//             }
//         }
//         if (n == 0) {
//             return;
//         }
//         lebelmap[lebels[index]] = lebel;
//         console.log(lebel, "  ", nums)
//         if (name.length == 1)
//             sets.push({ sets: name, size: n });
//         else
//             sets.push({ sets: name, size: n, label: "" });
//         index++;
//     }
//     ff([]);
//     console.log(sets);
//     let chart = venn.VennDiagram();
//     chart.wrap(false)
//         .width(640)
//         .height(640);
//     let div = d3.select("#venn").datum(sets).call(chart);
//     div.selectAll("text").style("fill", "white");
//     div.selectAll(".venn-circle path").style("fill-opacity", .6);
// }