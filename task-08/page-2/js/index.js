

function storage(){
    localStorage.setItem("lastname", "Kumar");
    alert("The value stored in the localStorage is: "+localStorage.getItem('lastname'))
}

function popup(){
    var time = new Date().toLocaleTimeString()
    var display = "The Current time is: " + time;
    window.alert(display);
}

function backgroundchanger(){
    document.body.style.background = "black" ;
}

function refresh(){
    window.location.reload(true)
}

function openWindow(){
    window.open("https://amfoss.in",'_blank');
}

function spam(){
    for(i=0;i<100;i++){
        console.log("I did it")
    }
}