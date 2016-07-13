/*
Package types provides primitives for creating types within a Phargo type graph.

The are a few core primitives provided.

The Node primitive is the core building block of a Phargo type graph. Most of
the data that your application will need to model will derive from the Node
primitive. A Node can be used to represent a blog post, a business, an address
or whatever standalone piece of data that you may wish to represent.  Nodes can
reference other Nodes to better structure data. For example, to represent a
business may want to store a name and an address. The name can be a property
directly on the business node, but you may wish to share the way you represent
an address across other data types. In this case, you would create a simple a
simple reference to an address type instead of directly storing the address
on the business node.

The Relationship primitive is closely related to a Node type, but provides extra
functinonality. A Relationship can model the relationship of one node to
another. That is, you may want to model a user's membership in an organization.
A Relationship would be used to create the connection between those two
entities, but also to store information about that relationship. For example,
you may want to store the user's role in the organization, like "Founder", or
their status in the organization, like "Paying Member", "Alumnus" or "Honorary
Member".

The Agent primitive is the final primitive available by default. An Agent is
the representation of some actor that has privileges to mutute data in the
Phargo type graph. An Agent might be a user or another application. Additional
functinonality exists around data that uses the Agent primitive, like
authentication and access policies.
*/
package types
