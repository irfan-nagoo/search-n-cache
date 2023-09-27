# search-n-cache

![image](https://github.com/raptor-23/search-n-cache/assets/142492599/c02fdc1c-a276-4395-9b55-07ed38268cc9)


This project covers some of the use cases for Searching and Caching aspects of an application or service. This project is packaged with search-n-cache-service which provides few basic blogging features and a valid use case for implementing Searching and Caching in GO (GOLANG) programming language. This project covers only real time indexing features of the underlying search technology.


## Searching

The Searching technology is generally used when we are expecting our database to support thousands or millions of record  for entities like customers, products, articles etc. and we want high performance discovery interface. If we are not expecting huge number of records in our database then in that case simple database reads could also suffice. The search technology in general have schemas which are called as collection/index. the collections/indexes are generally populated by using imports (Full import or Delta import) which are executed offline or using Realtime indexing APIs:

  1. Full imports configured on the search technology generally trigger a database SQL query to read data from the database and index it into the search technology. These imports are generally triggered by the scheduled job which is executed at specific interval. Therefore, the records are not available to search immediately when they are created in the database through this method.
  2. Delta import works in similar fashion however, accounts only those records which have changed since the last execution using last_modified_date.
  3. Realtime indexing makes sure that the record is indexed as soon as the record is created in the database. Hence, record is immediately available for searching. This project implements real time indexing.  

There are many search technologies available in the market like ElasticSearch, Solr, AWS OpenSearch (fork of ElasticSearch) etc. 

  1. **ElasticSearch** is based on Apache Lucene search APIs and provides lot of customization features and is generally suitable for big or enterprise level use cases in terms of records, systems using it etc. This project implements ElasticSearch interface.
  2. **Apache Solr** is also based on Apache Lucene search APIs however, provides less features as compared to ElasticSearch. Apache Solr is generally suitable for small search use cases used within an application. This project provides a place holder for Apache Solr interface however, implementation is not provided.
  3. **AWS OpenSearch** is a fork of ElasticSearch and hence provides features supported by standalone ElasticSearch on AWS Cloud.

## Caching 

The Caching technology is generally used for better performance reasons. The cache technology like Ehcache is generally within a standalone application however, it might not be suitable for distributed environment with clustering and load balancing aspects. In such environments, we have distributed caching technologies like Redis, Memcache etc.

  1. **Redis cache** (as a cluster) is generally used in distributed environment as in memory cache or data store. It has many other features like pub/sub to support various needs. This project implements Redis cache technology as a distributed cache.
  2. **Memcache** is also used as distributed in memory caching technology. This project provides a place holder for Memcache.   

## REST Interfaces with Blogging use case

This project has following REST interfaces:
  1. Article Search:
      - **Article List**   - This paginated APIs provides the list of articles sorted with latest articles at the top of the list. In the blogging UI app, the response of this api would be in the landing page.
      - **Article Search** - This paginated API returns the list of all the records matching the input search string with most probable record at the top. In the blogging UI app, this API would be triggered when the user tries to search any article.
  2. Article Detail:
      - **Get Article**    - This API return the article with the given ID. This API first tries to retrieve an article from cache, if  unavailable, reads database and sets record in cache, and returns the response. In the blogging UI app, this API will be invoked when the user clicks any article from the search/list response from above search APIs.
      - **Save/Update/Delete Article** - These APIs perform the specified action on the article and also syncs the record with the underlying search technology. In the blogging UI app, these APIs would support functionalities like Add blog, Edit Blog and delete a blog. 

This blogging Microservice could be extended by providing features like user login, user photo, adding documents, tags management etc. In the database, we might consider adding more tables like user, document (image and other docs), tags, category etc. with various cardinalities with main article table.


## Configuration

The Mysql DB, ElasticSearch, Redis need to be installed and should be up and running for this application to boot up and work properly. The default configuration should work smoothly

**Tech Stack:** Go 1.21.0+, gin-gonic, GORM, ElasticSearch, Redis, MySql

Here is the Swagger UI for various REST Endpoints: http://localhost:8080/swagger/index.html
