#!/usr/bin/env python

from flask import Flask, request, render_template, redirect, make_response


app = Flask(__name__)

@app.route('/notrust-check')
def auth():
    if "NTID" not in request.cookies:
        return '', 401
    return '', 200

@app.route('/notrust-login', methods=["GET", "POST"])
def login_do():
    if request.method == "GET":
        return render_template('login.html')

    user = request.form.get("username")
    password = request.form.get("password")

    if user == "admin" and password == "admin":
        response = make_response(
            redirect(request.cookies.get("NTREDIRECT"),
                                          code=302))
        response.set_cookie("NTID", "lolsegurojiji")
        return response
    else:
        return redirect("/notrust-login")

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=9000)
