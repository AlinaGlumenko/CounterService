let getCookie = (key) => {
    let arr = document.cookie.split(';').filter((item) => item.trim().startsWith(key+"="));
    let value = arr[0].split("=")[1];
    console.log("cookie: ", key, value);
    return value;
}

let setCookie = (name, value) => {
    if(document.cookie == "") {
        document.cookie = "max-age=60*60*24*365;"
    }
    document.cookie = name + "=" + value;
}

let cookieExistence = (key) => {
    if(document.cookie.split(';').filter((item) => item.trim().startsWith(key+"=")).length) {
        return true;
    }
    return false;
}

let getCount = () => {
    let params = {};
    let link = window.location.href;
    if(!cookieExistence(link)) {
        setCookie(link, 'true');
        params = {"pagename": link, "visit": "true"}
    } else {
        params = {"pagename": link, "visit": "false"}
    }
    let url = new URL("http://localhost:8000/api/count_users");
    Object.keys(params).forEach(key => url.searchParams.append(key, params[key]));
    let el = document.querySelector("#counter");
    el.innerHTML = 0;    
    (async() => {
        try {
          var response = await fetch(url);
          var data = await response.json();
          el.innerHTML = data;
        } catch (e) {
          console.log(e);
        }
    })();
};

document.addEventListener("DOMContentLoaded", getCount);