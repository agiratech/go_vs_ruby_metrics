# Golang vs Ruby Performance Metrics

It all starts with a Rails application as usual. In one particular situation, we need to upload complex csv record to ETL DB for reporting purpose with over 500k entries for every kickstart.

The initial code was written in Ruby and run via sidekiq jobs but was taking more time to finish the job.

This was a bottleneck we decided to eliminate and since there was a lot of scope for concurrency here, we decided to use Golang for this.

we @agiratech wondered while i'm executing same process in Golang, the numbers are surprised me , we shared our experience here. To ensure the metrics, we ran the same process in c4.xlarge and c4.2xlarge in amazon web services.

From our study, the overall the performance of Golang is around 5x faster than Ruby.

You can refer golang code in master branch and ruby code in rails branch.

For more graph metrics, you can refer this link (http://www.agiratech.com)

To hire golang developers contact http://www.agiratech.com/golang-development