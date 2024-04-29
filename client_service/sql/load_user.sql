SELECT name, surname, birthdate, email, phone_number
FROM users
WHERE login=$1
