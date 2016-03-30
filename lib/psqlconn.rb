module Psqlconn
  class Psql

    TABLE_NAME = "ruby_csv"

    def initialize
      @conn = PG::Connection.open(:dbname => 'ruby_test_csv', user: 'postgres', password: 'root')
    end

    def create_table(header)
      check_table = "CREATE TABLE IF NOT EXISTS #{TABLE_NAME}"
      set_primary_key = "(ID integer PRIMARY KEY,"
      set_variables = ""
      header.to_a.each_with_index do |h, i|
        value = header.length == i+1 ? "#{h} TEXT)" :  "#{h} TEXT,"
        set_variables += value
      end
      final_query = check_table+set_primary_key+set_variables
      execute_query(final_query)
    end

    def insert_record(data)
      query = "INSERT INTO #{TABLE_NAME} values(#{data})"
      execute_query(query)
    end

    def truncate_data
      query = "TRUNCATE TABLE  #{TABLE_NAME}"
      execute_query(query)
    end

    def execute_query(query)
      @conn.exec_params(query)
    end
  end
end