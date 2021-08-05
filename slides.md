---
theme: seriph
background: https://unsplash.com/collections/2476111/1920x1080
class: text-center
highlighter: shiki
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

  Learn more at [Sli.dev](https://sli.dev)
title: Software design and refactoring
---

# Talk software design and refactoring

## Phan Minh Tâm

tam.phan1@tiki.vn
---

# Content
- Recap recap Grokking Techtalk #37: Software design and refactoring
	-	OPP is not about and inheritance.
 	-	Sofware design is about cracking complexity.
	-	reference:
	 	-   https://www.youtube.com/watch?v=7PGCvpJl_0o
     	-   https://www.slideshare.net/GrokkingVN/software-design-and-refactoring-215436212
    
- Case study: 
	-	Kafka consumer
  	-	Software scalable


---

# OPP Software design ?
- OOP everyone know it - I won't explain it. it already includes SOLID
- SOLID review
  - The Single-responsibility principle: "There should never be more than one reason for a class to change." In other words, every class should have only one responsibility.
  - The Open–closed principle: "Software entities ... should be open for extension, but closed for modification."
  - The Liskov substitution principle: "Functions that use pointers or references to base classes must be able to use objects of derived classes without knowing it." See also design by contract.
  - The Interface segregation principle: "Many client-specific interfaces are better than one general-purpose interface."
  - The Dependency inversion principle: "Depend upon abstractions, not concretions"
- https://en.wikipedia.org/wiki/SOLID

<br>
<br>


<style>
h1 {
  background-color: #2B90B6;
  background-image: linear-gradient(45deg, #4EC5D4 10%, #146b8c 20%);
  background-size: 100%;
  -webkit-background-clip: text;
  -moz-background-clip: text;
  -webkit-text-fill-color: transparent; 
  -moz-text-fill-color: transparent;
}
</style>


---
layout: two-cols
---

<template v-slot:default>

# Functional programming
```go
Router.Routing(HandlerFn, "endpoint") {
    return HandlerFn(ServiceFn) {
      return ServiceFn(repositoriesFn...){
        repositoriesFn[0]() -> models.A
        repositoriesFn[1]() -> models.B
        ...
        dtos = combine(models.A, ...) 
        return dtos 
      }()
    }() 
  }
```

<v-click>

Functional programming
<ul>
  <li>Code is short</li>
  <li>Meaning full</li>
</ul>

</v-click>

<v-click>

But
<ul>
  <li>Is easy maintain ?</li>
  <li>How to deal complexity ?</li>
</ul>

</v-click>

</template>
<template v-slot:right>

# OOP
```go
Router.Routing(Handler.HandlerFn, "endpoint") 
Handler.HandlerFn() {
  return serviceA.DoSomething()
}
serviceA.DoSomething() {
  // handle business logic
  models: A = repository1.GetSomeThing()
  models: B = repository2.GetSomeThing()
  ...
  return combine(A, B) 
}
```
<v-click>

OPP
<ul>
  <li>Easy maintain</li>
  <li>Easy open for extension</li>
  <li>Complexity encapsulation</li>
</ul>

</v-click>

<v-click>

But are you sure ???

</v-click>
</template>

---
layout: two-cols
---

<template v-slot:default>

## Example about Inheritance
### use case
<v-click>

```js
class BaseHandler {
  // property
  service: serviceA

  func method1()
  func method2()
  ......
  abstract func DoSomething() // fake method
  func RealDoSomething() {
    method1()
    DoSomething()
    method2()
  }
}

class ProductHandler(BaseHandler) {
  func DoSomething() {
    //... implementation
  }
}

```

</v-click>

</template>
<template v-slot:right>

## Is this complex ?

<v-click>

- Ok! single class isn't complex right !!

</v-click>

<v-click>

- Is it meet the SOLID conditions ?

</v-click>

<v-click>

- Something like that

```js
class BaseService{
  ...
}

class ServiceA(BaseService){
  // property
  repo1: repository1
  repo2: repository2
  ...
}
class Baserepository{
  ...
}
class Repository1(Baserepository){
  ...
}
```

</v-click>

<v-click>

- What do you think ?

</v-click>

</template>

---
layout: two-cols
---

<template v-slot:default>

## Example about Inheritance
### Refactor
<v-click>

```js
interface Worker {
  DoSomething()
}

class ProductWorker implement Worker {
  DoSomething() {
    //.....
  }
}

class ProductHandler {
  worker      : ProductWorker // composite pattern
  service: serviceA

  func method1()
  func method2()

  func DoSomething() {
    method1()
    worker.DoSomething()
    method2()
  }
}
```

</v-click>

</template>

<template v-slot:right>

## Refactor

<v-click>

- Do same thing for class Service, Repository

</v-click>

<v-click>

- It meet the SOLID conditions 

</v-click>

<v-click>

- What do you think ?

</v-click>

<v-click>

```js
class BaseHandler {
  func method1() {

  }
  func method2() {

  }
}
class ProductHandler implement Worker {
  baseHandler: BaseHandler
  DoSomething() {
    baseHandler.method1()
    //....
    baseHandler.method2()
  }
}

```

</v-click>

</template>

---

# Inheritance multi level  

```python
class KafkaSpider(scrapy.Spider):
  site_id = id
  """Do a lot of logic kafka consumer, cookies manager ...."""
  class Meta:

class KafkaProductDetailSpider(KafkaSpider):
  """Do logic process kafka message, send kafka message"""

class BHXProductDetailSpider(KafkaProductDetailSpider):
  """Do logic parse website"""
```
<v-click>

- This code is makes sense ? Yes, No

</v-click>

<v-click>

- Why `class Meta` is not composite ?

</v-click>

<v-click>

- How do cracking complexity ?

</v-click>

---

# How do cracking complexity?

```python
class Consumer():
  site_id = id
  def __init__(SpiderFactory, SiteParserFactory, producer, ...):
    ... settup consumer
    site_parser = SiteParserFactory(site_id) # do logic how to create Parser adapt with site_id
    self.spider = SpiderFactory(site_id, site_parser) # do logic how to create Spider adapt with site_id
    self.producer = Producer(site_id) # which topic send msg
  
  def consume(self):
    messages = self.kafka_message_consumer.poll(5)
    process_kafka_message(messages)
    commit_msg()

  def process_kafka_message(self, message):
    self.spider.crawl(msg)
    producer.send_msg(...) # factory msg adapt with logic
  ....
  self.spider.signals.connect(self.consume, signal=signals.spider_idle)
  self.spider.signals.connect(self.consume, signal=signals.item_scraped)
```

---
layout: two-cols
---

<template v-slot:default>

## How to design software easy refactor or scalable
<br/>
<v-click>

- Dependent class to interface not instance

</v-click>
<v-click>

- Two class know interface not implementation

</v-click>
<v-click>

- Not dependent in framework specific
  
</v-click>
  
<v-click>

- Not dependent in database specific
  
</v-click>

<v-click>

- Use design patterns effective

</v-click>

<v-click>

  - * 4 group: Creational patterns, Structural patterns, Behavioural patterns, Concurrency patterns
  - * Think use together
  - * Some patterns can be use single Factory, Builder, Singleton ...

</v-click>
  
</template>

<template v-slot:right>

## Scalable
<br/>
<br/>
<br/>
<v-click>

## what things affect scalable ?

</v-click>

<v-click>

- Programing use develop

</v-click>

<v-click>

- Coding(code size, code base ...)

</v-click>

<v-click>

- Data size

</v-click>

<v-click>

- Datababse

</v-click>

<v-click>

- System design

</v-click>

</template>

---

# [Programing use develop](https://github.com/mtamphan/my-talk.io/tree/main/benchmark)

## python is too slow
### Benchmark python vs java vs golang: problem Matrix Multiplication (Demo)

Chip:	Apple M1
Total Number of Cores:	8 (4 performance and 4 efficiency)
Memory:	16 GB

<v-click>

|             | python    | java     | golang    | golangv2 |
|-------------|-----------|----------|-----------|----------|
| 256 x 256   | 4.10 s    | 48 ms    | 62.85 ms  | 250 ns   |
| 512 x 512   | 31.47 s   | 399 ms   | 409.17 ms | 295 ns   |
| 1024 x 1024 | 4:49.66 m | 3.301 s  | 3.27 s    | 334 ns   |
| 2048 x 2048 | too slow  | 12.845 s | 25.95 s   | 375 ns   |

</v-click> 

- Python slower java: ~ 10 times
- Python slower golang: ~ 10 times
- Golang vs Java: same
- Learn and use effective one programming is important.


---
layout: two-cols
---

<template v-slot:default>
  
# Coding
<br/>
<v-click>

- Think monolith first not microservices first
	- I never see someone design microservices first correctly unless 10 years experience in software.
  
</v-click>

<v-click>

- Coding think how to split it to service

</v-click>

<v-click>

- Coding thinks how to parallelize and effectively use the programming language in use

</v-click>

<v-click>

- use effectively multi Threading language support
- [python](https://en.wikibooks.org/wiki/Python_Programming/Threading)
- [java](https://docs.oracle.com/javase/8/docs/api/java/lang/Thread.html)
- [golang](https://golang.org/doc/effective_go)

</v-click>
  
</template>

<template v-slot:right>

## Example
<br/>
<v-click>

### Product update from seller center:
  - update tiki product info, update price ..., update index
  
</v-click>
  
<v-click>

```go
func process_celler_center_event(product) {
  doUpdateTikiProductInfo(product)
  doUpdateTikiProductPrice(product)
  doUpdateIndexProduct(product)
}
```
  
</v-click>

<v-click>

```go
func process_celler_center_event(product) {
  threadpool <- doUpdateTikiProductInfo(product)
  threadpool <- doUpdateTikiProductPrice(product)
  threadpool <- doUpdateIndexProduct(product)
}
```
  
</v-click>

<v-click>

```go
func process_celler_center_event(product) {
  for observer in observers:
    observer.DoProcessEvent(product)
}
```
  
</v-click>

</template>

---
layout: two-cols
---

<template v-slot:default>

# Using observer pattern

<v-click>

```go
observer = make([]observer, 0)
func process_celler_center_event(product) {
  for observer in observers:
    observer.DoProcessEvent(product)
}
```
  
</v-click>
  
<v-click>

```go
ProductInfo
func doProcessEvent(product) {
  DoUpdateInfo()
  DoUpdatePrice()
}
```
  
</v-click>
<v-click>

```go
ProductIndex
func doProcessEvent(product) {
  DoIndex()
}
```
  
</v-click>
  
</template>

<template v-slot:right>
  
<v-click>
 
 # Split 2 consumer group 

</v-click>
   
<v-click>

```go
consumer ProductIndex
func doProcessEvent(product) {
  DoIndex()
}
```
  
</v-click>
  
<v-click>

```go
consumer ProductInfo
func doProcessEvent(product) {
  DoUpdateInfo()
  DoUpdatePrice()
}
```
  
</v-click>

<v-click>
 
 ## add queue process

</v-click>

<v-click>

```go
consumer ProductInfo
func doProcessEvent(product) {
  DoUpdateInfo()
  pushEventPriceChange()
}
```
  
</v-click>

<v-click>

```go
consumer PriceChange
func doProcessEvent(product) {
 doPriceChange()
}
```
  
</v-click>  
  
</template>


---

# Database

### Recap Grokking TechTalk #20: PostgreSQL Internals 101
[PostgreSQL Internals 101](https://www.slideshare.net/GrokkingVN/grokking-techtalk-20-postgresql-internals-101)
## Use index correctly
<br/>
<v-click>
  
- Column index

</v-click>

<v-click>
  
- B+ Tree vs Hash
 
- - [MySql document](https://dev.mysql.com/doc/refman/8.0/en/index-btree-hash.html)
- - [stack overflow](https://stackoverflow.com/questions/7306316/b-tree-vs-hash-table)

</v-click>

<v-click>
  
- fulltext search, use indexer libary
- [lucene](https://lucene.apache.org/core/), [solr](https://solr.apache.org/guide/6_6/introduction-to-solr-indexing.html#:~:text=A%20Solr%20index%20can%20accept,as%20Microsoft%20Word%20or%20PDF.)

</v-click>

---
layout: two-cols
---

<template v-slot:default>
  
# Database

### Current issue tiki_product_info
<br/>

<img src="/tiki_p1.png" class="h-55 rounded shadow" />
<img src="/tiki_p2.png" class="h-40 rounded shadow" />

</template>
<template v-slot:right>
 
# _
  
### seller_id = 1 is 1P product
 
### sql query pricing tool
<v-click>

```sql
  SELECT * FORM tiki_product_info 
  WHERE seller_id = 1 AND .... filter condition
```

</v-click>
  
### Root cause it too slow

<v-click>

- fullscan stable tiki_product_info

</v-click>

<v-click>

- seller_id not unique using index not effect performent query
  
</v-click>

### Solution

<v-click>

- use partion table ?

</v-click>
<v-click>

- - only show 3P product or show all ?

</v-click>

### My Solution

<v-click>

- index custom column

</v-click>
<v-click>

Use custom id, 1P product id <= 5000000 (5tr), 3P > 5000000 (5tr). Index this column

</v-click>

</template>


---

# Database

## Embedded database
<v-click>
  
- [Berkeley DB](https://www.oracle.com/database/technologies/related/berkeleydb.html#:~:text=Berkeley%20DB%20is%20a%20family,for%20data%20access%20and%20management.)
  
</v-click> 
  
<v-click>
  
- [Kyoto cabinet](https://dbmx.net/kyotocabinet/)
  
</v-click>
  
<v-click>
  
- [rocksdb](https://rocksdb.org/)
  
</v-click>

## Graph data base
<br/>
<v-click>
  
- [neo4j](https://neo4j.com/) 
  
</v-click>

<v-click>
  
- [JanusGraph](https://cloud.google.com/architecture/running-janusgraph-with-bigtable) - alow customize
  
</v-click>

## how to build time series database ?
<br/>
<v-click>
  
- Use key-value database
  
</v-click>

<v-click>
  
- different value with two point 
  
</v-click>

---
layout: image-right
image: ./IMG_0140.jpeg
---

# System design

### problem

- We are going to develop a commercial website for selling products where:

- Each product has a number of available items in stock. The system should be able to process at least N = 6000 concurrent requests for viewing or purchasing products. The system is only allowed to have at most S = 6 servers, where up to 3 servers can be used as relational databases. One relational database can serve at most C = 300 concurrent connections.

- The system must: (main requirements) ensure data consistency.


---
layout: two-cols
---

<template v-slot:default>

# Handle concurrency and race condition

### concurrency algorithms

<div v-click="1">

- Read write lock (easy implement)

</div>

<div v-click="2">

- Time stamp (advanced technical)

</div>

<div v-click="4">

- How to implement `read_ts`, `write_ts`.

</div>

<div v-click="5">

- synchronized update `read_ts`, `write_ts`.

</div>

<div v-click="6">

- Use Atomic i64 ???

</div>


</template>

<template v-slot:right>

### Time stamp (advanced technical)

<div v-click="3">

```go
func Read(context, A, ts) {
  if ts > write_ts {
    read(a)
    update(read_ts = current_time)
    if update == false {
      context.rollback()
    }
  } else {
    context.rollback()
    retry()
  }
}
func Write(context, A, ts) {
  if ts > read_ts {
    if ts > write_ts {
      write(a)
      update(write_ts = current_time)
      if update == false {
        context.rollback()
      }
    }
    // inogre
  } else {
    context.rollback()
    retry()
  }
}
```

</div>
</template>
