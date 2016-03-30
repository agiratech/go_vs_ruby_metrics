class HomeController < ApplicationController
  require 'csv'

  def index

  end

  def import
    start_time = Time.now
    array_header = ["policy_id",
                   "state_code",
                   "county",
                   "eq_site_limit",
                   "hu_site_limit",
                   "fl_site_limit",
                   "fr_site_limit",
                   "tiv_2011",
                   "tiv_2012",
                   "eq_site_deductible",
                   "hu_site_deductible",
                   "fl_site_deductible",
                   "fr_site_deductible",
                   "point_latitude",
                   "point_longitude",
                   "line",
                   "construction",
                   "point_granularity"
                   ]

    CSV.read('./db/samples/myfile_sample.csv', { col_sep: "," }).each_slice(200).each_with_index do |rows,i|
      array_hash =[]
      rows.each_with_index { |row,j| next if (i==0 && j==0);array_hash << Hash[[array_header,row].transpose] }
      RubyCsv.create(array_hash)
    end

    end_time = Time.now
    render :json => {status: 200, msg: "success", process_time: end_time - start_time}
  end

  def search_api
    start_time = Time.now
    results = RubyCsv.where("county iLIKE ?", "%#{params[:county]}%")
    end_time = Time.now
    render :json => {status: 200, msg: "success", process_time: end_time - start_time, record:{ count: results.count, data: results} }
  end

end
