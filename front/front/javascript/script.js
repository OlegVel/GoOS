function myFunction() {
  document.getElementById('right').textContent='hello';
}

function jsonFunc() {
  var xhr = new XMLHttpRequest();
  var url = 'http://localhost:5000/api';

  xhr.open("POST", url, true);
  xhr.setRequestHeader("Content-Type", "application/json");
  xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
      var json = JSON.parse(xhr.responseText);
      console.log(json.test + ", " + json.password);
    }
  };
  var data = JSON.stringify({"test": "hey@mail.com", "password": "101010"});
  xhr.send(data);
}
