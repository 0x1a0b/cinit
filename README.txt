CONTAINER INIT

DESIGN
    The general consensus in the container operations community is that a
    deployed container should have only a single process running in it. This
    "best practice" is slightly loosened for applications that fork/spawn to
    create a cluster (see NodeJS cluster module, python's gunicorn, etc) but
    that appears to be the extent of it.

    I disagree with this assessment. Muti-process applications are extremely
    common in practice, however the basis for process separations is typically
    tied to privileges and framed around security. The concept of a
    heterogeneous multiprocess application isn't given much thought. I find it
    quite useful to bundle my containers at the level of discrete applications.
    
    This causes many challenges for deploying medium sized systems with
    containers. Using a single process per container forces the operator to use
    -- typically quite heavyweight -- container orchestration solutions such as
    Amazon ECS or Kubernetes. Adding pods/services simply adds another layer of
    configuration to an otherwise simple system. Not to mention the overhead of
    keeping the container scheduler running in the first place.
    
    The other option is to treat a container like a virtual machine, running an
    init program as if it were a full operating system. This is also undesirable
    for a few reasons. Firstly it will result in running may extraneous
    processes that cloud applications don't typically need require. Either this
    or run a highly customized version the OS in question which has its own
    difficulties. The other reason is that the process a docker container runs
    does not have the same semantics as an init process run by a kernel. The
    duties expected of the process as different and signals send/received and
    their meanings are not the same [1]. This makes traditional process managers
    (which is the vast majority them) ill suited to this task.

    Container init is my attempt to land somewhere in the middle. Designed from
    the ground up to multiplex processes written by different teams while
    handling the semantics  expected from a docker containers entry point.


BUILDING
    $ make all



[1] TODO: Actually figure out where you read this

