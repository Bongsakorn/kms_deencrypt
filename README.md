# KMS Encrypt-Decrypt

This command use for decrypt text that decrypted with aws kms.

### Usage

To decrypt message

```terminal
❯ ./kms_decryption decrypt --help
Usage of decrypt:
  -profile string
    	aws credential profile. (optional) (default "default")
  -text string
    	Decrypted text. (Required)
```

To encrypt message

```terminal
❯ ./kms_decryption decrypt --profile dev-plugin --text NP+FAwEBB3BheWxvYWQB/4YAAQMBA0tleQEKAAEFTm9uY2UB/4gAAQdNZXNzYWdlAQoAAAAZ/4cBAQEJWzI0XXVpbnQ4Af+IAAEGATAAAP4Ddf+GAf+oAQIDAHjNNdp6EVDb/UpYNTKgFDr2kecnlcBIXypCrKN/OPYvVQG9c6FBOnUnx/Hb4SzQXKnyAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMfLAWdSD+XOtDKQD0AgEQgCvCWt/3pritVmZ51cBTCqvWdCBUYFWLyzlYN10Yc8A0CPUZI8LSLOGlzVvEARj/2f+A/+r/1//LJ2v/7//b//T/vyMW/7z/y/+///UeY//GFP/WUP+xAf4CmV+gXYMaW/hnht2lAyue3EGM64HchHDN5JpB8o/r6N2CCtV3lMOl7Tz3M0SOMIi19RyFUyH1nI08kPz0DCXpArWxaCcIP4fYXoy7xgdB2v2lRqt9BHV4skkmN5NmYwK04pDQTz6PcVn4YGKn7gUumrBTsM4ICzOsH9UQ85S/CFRYQlA9nYuGL3qS2qgKvyJXxev4WShskwlimt0rpZOQrlNCyKQQT6mQdYUcALTU3zzqaA4YG0tJjGeushPLfNww7q5aoXGADJm3GF3HKSgMf/8CI2NO2n8Z2isFs9if73XFlVhyWgDMN5Wkwv7LWGi4cu7VZRDpbCJHoxIVB1iYQwV0JceREWFz8QUs6reRd9kq2G6EQQX+jBZj+DhR1x7jsbDrx9lVNyUCMAhcVpz4VxuiuXp9raifl7dul/0DFcYBD3UKGvAsx0Ks7yX4kQchbfgtAbQgzc/lWfUFcgLy1JZPuk/WJCq5BwNkt33Qx5WlB7NSSHZEZ1S+AeAwgThlp4DL8Wywt07ZWWb5djazH2Rl8msT/FbzTvyPO46RSUxHMG+ocmvsNqVsfLRjszD7xskAZ25leQZnVdywxpO+E6L3Rf+QYLM/z5D+IQXo3gLCZE59hT9fM84gFCbxCenc0dWKYKIAiv4bMXAl7uRxxK6VHNnEurWYPsnm3069Lox3/fAc5jIi8gi0ZaP4CpBs5ztgepGcR+KOCLteQ94c5Y9GcRaW3GbQ740/ZUs4cIhRMS8mfGsXLFuGehv2yblNll9GzJB2zLwuqrGsbemXNE9C6IERGi7//Ps5sPrWevwO6ob2PYLpHWoTflHFVFNiEhz+1MmDnd5FfHw1K8jSJjZeInEn9dYkhCP9aKbmWbLmfCaw9fP9aikZAA==
=====================
{"package_id":1,"assurer":{"email":"Test@test.com","phone_number":"0950073000","id_card":"1046176319089","passport_id":"","title":"นางสาว","firstname_th":"เทส","lastname_th":"ห้า","date_of_birth":"1976-08-24T17:00:00Z","address":"","house_no":"11/1","moo":"1","address_name":"เทสเพลส","soi":"รัชดา18","road":"รัชดา","sub_district":"ห้วยขวาง","district":"เขตห้วยขวาง","province":"กรุงเทพมหานคร","post_code":"10310"},"coverage_from":"2020-12-25T17:00:00Z","coverage_to":"2021-12-25T17:00:00Z","consent_1":false,"consent_2":false}
```

### Example

Decryption

```terminal
❯ ./kms_deencrypt decrypt --profile plugin-prod --text 'NP+BAwEBB3BheWxvYWQB/4IAAQMBA0tleQEKAAEFTm9uY2UB/4QAAQdNZXNzYWdlAQoAAAAZ/4MBAQEJWzI0XXVpbnQ4Af+EAAEGATAAAP4Ddv+CAf+oAQIDAHgFcYo+WZaEjT/8FPcS02QtpuxP6kvauIiFljAGdiRq9QHzGLdMu1DF8s4PjHjgUrIeAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMW5XbRGzq2qdvM8cMAgEQgCtVMN6k0MtLJE96GvZvbI5E/rBprFSifwhNjv2nuxEdKX6azGPKa0544xZeARg1dUdJ/+4l/7Vv/8j/mWv/4P/C/6j/ilr/zv/yTDj/ky3/+W4B/gKeB4v1HgB5ZfW5mHoxonsAUMDjVAj4MyhAbpYhXrLIDVzEgDDp7nrx84r9NItyrmP/2VYJ1k4y4ZM01q/VG1u2DSHj2MWrOfBYbKJiceXXBkgH5/WXBHnXPm4RjKKpYWqIdDxwlZbVd3Opmps1le+IdHbqJt4BRsdukvjAIV8cgyTxxG9nhnTbvPWJ1+1QWLvkO2Jh+XwW/3oNRmmiT24EHu2thABDeTjKyqE43TlpDNtn4MRfaa9qKznfvSZr6sXqVIgqsE3a0qsot3cABGCU6On85FUe9k9DeXE+5k2lFu5Tki+xFGiXUmSMFIQuCIky49yuDyt7bT4RK5pPbXocLTgDFecxS7j8NXck8kogm/7s5/B13yYzR6lHPHFbrXWesKt5JurPuwxgNM7oF98roNaAADKxn82Upc4mTXk0gEZwyMEnKAC3SCFjVibcwdUazHSyZysRhbEZD9JAASea3kC0zws6ruqszY7Vk4ml/4kEJuRtf3Oh6eorVVn6O+tdozGLfIbaRQq3CS/ElpZFzx5vRDdpTvDz/x9q5fI/TNejscyjpRflytvUGGIGLct1SjZ60DZfZ/qIAh5rymEd56KzUsNkDDbvIdoftiR1JyK4WpUPVAL1Memsjru0Kt7PYhUwP24YdosoQsJaqoQun0VmpwjMU3xAE6n54C1LbxgN7NM1+ne2tJsfdyMJoCMM0fnrG+YMQyPK3p1RqmiW6c3q44G3gFBskVq+FNobydBXHCYIx9Se6LkiZk59o80x7NcVEBwJufXfo09e2vHj2f6kIFcVzu1qOqpq6Qwx7jDElRf/NHZ7G+PkfgNiqXEUnN4D4fZQ4K/i/exMgTARHjhgcduBdY3dTiGqUxw6lYJ6PGelHUWBxL6seJXw+AA='
=====================
{"package_id":2,"assurer":{"email":"tawan.vangvasu@gmail.com","phone_number":"0950073000","id_card":"3101401031319","passport_id":"","title":"นางสาว","firstname_th":"ตะวัน","lastname_th":"วังวสุ","date_of_birth":"1982-08-24T17:00:00Z","address":"","house_no":"93/911","moo":"","address_name":"ไอดิโอ สุขุมวิท93","soi":"","road":"สุขุมวิท","sub_district":"บางจาก","district":"พระโขนง","province":"กรุงเทพ","post_code":"10260"},"coverage_from":"2020-12-25T17:00:00Z","coverage_to":"2021-12-25T17:00:00Z","consent_1":true,"consent_2":true}
```

Encryption

```terminal
❯ ./kms_deencrypt encrypt --profile plugin-prod --keyname LoggingKey --text '{"package_id":2,"assurer":{"email":"tawan.vangvasu@gmail.com","phone_number":"0950073000","id_card":"3101401031319","passport_id":"","title":"นางสาว","firstname_th":"ตะวัน","lastname_th":"วังวสุ","date_of_birth":"1982-08-24T17:00:00Z","address":"","house_no":"93/911","moo":"","address_name":"ไอดิโอ สุขุมวิท93","soi":"","road":"สุขุมวิท","sub_district":"บางจาก","district":"พระโขนง","province":"กรุงเทพ","post_code":"10260"},"coverage_from":"2020-12-25T17:00:00Z","coverage_to":"2021-12-25T17:00:00Z","consent_1":true,"consent_2":true}'
=====================
NP+BAwEBB3BheWxvYWQB/4IAAQMBA0tleQEKAAEFTm9uY2UB/4QAAQdNZXNzYWdlAQoAAAAZ/4MBAQEJWzI0XXVpbnQ4Af+EAAEGATAAAP4Ddv+CAf+oAQIDAHgFcYo+WZaEjT/8FPcS02QtpuxP6kvauIiFljAGdiRq9QHzGLdMu1DF8s4PjHjgUrIeAAAAbjBsBgkqhkiG9w0BBwagXzBdAgEAMFgGCSqGSIb3DQEHATAeBglghkgBZQMEAS4wEQQMW5XbRGzq2qdvM8cMAgEQgCtVMN6k0MtLJE96GvZvbI5E/rBprFSifwhNjv2nuxEdKX6azGPKa0544xZeARg1dUdJ/+4l/7Vv/8j/mWv/4P/C/6j/ilr/zv/yTDj/ky3/+W4B/gKeB4v1HgB5ZfW5mHoxonsAUMDjVAj4MyhAbpYhXrLIDVzEgDDp7nrx84r9NItyrmP/2VYJ1k4y4ZM01q/VG1u2DSHj2MWrOfBYbKJiceXXBkgH5/WXBHnXPm4RjKKpYWqIdDxwlZbVd3Opmps1le+IdHbqJt4BRsdukvjAIV8cgyTxxG9nhnTbvPWJ1+1QWLvkO2Jh+XwW/3oNRmmiT24EHu2thABDeTjKyqE43TlpDNtn4MRfaa9qKznfvSZr6sXqVIgqsE3a0qsot3cABGCU6On85FUe9k9DeXE+5k2lFu5Tki+xFGiXUmSMFIQuCIky49yuDyt7bT4RK5pPbXocLTgDFecxS7j8NXck8kogm/7s5/B13yYzR6lHPHFbrXWesKt5JurPuwxgNM7oF98roNaAADKxn82Upc4mTXk0gEZwyMEnKAC3SCFjVibcwdUazHSyZysRhbEZD9JAASea3kC0zws6ruqszY7Vk4ml/4kEJuRtf3Oh6eorVVn6O+tdozGLfIbaRQq3CS/ElpZFzx5vRDdpTvDz/x9q5fI/TNejscyjpRflytvUGGIGLct1SjZ60DZfZ/qIAh5rymEd56KzUsNkDDbvIdoftiR1JyK4WpUPVAL1Memsjru0Kt7PYhUwP24YdosoQsJaqoQun0VmpwjMU3xAE6n54C1LbxgN7NM1+ne2tJsfdyMJoCMM0fnrG+YMQyPK3p1RqmiW6c3q44G3gFBskVq+FNobydBXHCYIx9Se6LkiZk59o80x7NcVEBwJufXfo09e2vHj2f6kIFcVzu1qOqpq6Qwx7jDElRf/NHZ7G+PkfgNiqXEUnN4D4fZQ4K/i/exMgTARHjhgcduBdY3dTiGqUxw6lYJ6PGelHUWBxL6seJXw+AA=
```

That's it, peace :v:
