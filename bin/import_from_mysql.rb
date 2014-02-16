#!/usr/bin/env ruby

require 'mysql2'
require 'sqlite3'

ROWS = %w{id month day title excerpt source content summary}

SQL = <<END_SQL
  INSERT INTO
    jft
    (#{ROWS.join(', ')})
  VALUES
    (#{Array.new(ROWS.count, '?').join(', ')})
END_SQL

mysql = Mysql2::Client.new(host: 'localhost', username: 'root', database: 'nevercured')
sqlite = SQLite3::Database.new(File.expand_path(File.join(
  File.dirname(__FILE__), '..', 'db', 'nevercured.db')))

mysql.query('SELECT * FROM jft').each do |row|
  puts "Adding: #{row['id']}"
  values = ROWS.map { |column| row[column] }
  sqlite.execute(SQL, *values)
end
