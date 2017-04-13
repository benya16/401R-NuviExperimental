import tensorflow as tf


batch_size = 1
hidden_units = 40
n_input = batch_size #  dta input (img shape: 28*28)
n_classes = 2 # total classes (0-1 digits)
dropout = 0.75 # Dropout, probability to keep units
hidden = .5
num_columns = 4

model_path = "C:/Users/Greg/Anaconda3/envs/tensorflow/Scripts/model.ckpt"
  

x = tf.placeholder(tf.float32, [n_input,num_columns])
y = tf.placeholder(tf.float32, [n_input, n_classes])
keep_prob = tf.placeholder(tf.float32) #dropout (keep probability)
keep_hidden = tf.placeholder(tf.float32)


def init_weights(shape):
    return tf.Variable(tf.random_normal(shape, stddev=0.01))


def model(x,  w_h, w_h2, w_o, keep_prob,keep_hidden):
    x = tf.nn.dropout(x, keep_prob)
    h = tf.nn.relu(tf.matmul(x, w_h))
    h = tf.nn.dropout(h, keep_hidden)
    h2 = tf.nn.relu(tf.matmul(h, w_h2))
    h2 = tf.nn.dropout(h2, keep_hidden)
    return tf.matmul(h2, w_o)
	
# Store layers weight & bias
w_h = init_weights([num_columns, hidden_units])
w_h2 = init_weights([hidden_units,n_classes])
w_o = init_weights([n_classes, n_classes])


# Construct model
pred = model(x, w_h, w_h2, w_o, keep_prob, keep_hidden)

# Define loss and optimizer
#cost = tf.reduce_mean(tf.nn.softmax_cross_entropy_with_logits(logits=pred, labels=y))
#optimizer = tf.train.AdamOptimizer(learning_rate=learning_rate).minimize(cost)

# Evaluate model
#correct_pred = tf.equal(tf.argmax(pred, 1), tf.argmax(y, 1))
#accuracy = tf.reduce_mean(tf.cast(correct_pred, tf.float32))

#get prediction
prediction =  tf.argmax(pred,1)

init = tf.global_variables_initializer()

# Later, launch the model, use the saver to restore variables from disk, and
# do some work with the model.
with tf.Session() as sess:
    # Restore the model
    tf_saver = tf.train.Saver()
    tf_saver.restore(sess, "C:/Users/Greg/Anaconda3/envs/tensorflow/Scripts/model.ckpt")
	tf.run(prediction,x) 


##print_tensors_in_checkpoint_file("C:/Users/Greg/Anaconda3/envs/tensorflow/Scripts/model.ckpt", tensor_name ="" , all_tensors = True)
  
